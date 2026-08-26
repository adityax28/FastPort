using System.ComponentModel.DataAnnotations;
using Aditya.Portfolio.Models;
using Aditya.Portfolio.Services;
using Microsoft.Extensions.Options;

var builder = WebApplication.CreateBuilder(args);

// Render (and most PaaS hosts) inject PORT
var port = Environment.GetEnvironmentVariable("PORT");
if (!string.IsNullOrWhiteSpace(port))
{
    builder.WebHost.UseUrls($"http://0.0.0.0:{port}");
}

builder.Services.Configure<PortfolioOptions>(
    builder.Configuration.GetSection(PortfolioOptions.SectionName));
builder.Services.AddSingleton<PortfolioData>();
builder.Services.AddSingleton<ContactStore>();

var app = builder.Build();

app.UseDefaultFiles();
app.UseStaticFiles();

var api = app.MapGroup("/api");

api.MapGet("/profile", (PortfolioData data) => Results.Ok(data.GetProfile()));

api.MapGet("/status", (PortfolioData data, IHostEnvironment env) =>
    Results.Ok(data.GetStatus(env)));

api.MapPost("/contact", (ContactRequest request, ContactStore store, IOptions<PortfolioOptions> options) =>
{
    var errors = new List<string>();

    if (string.IsNullOrWhiteSpace(request.Name) || request.Name.Trim().Length < 2)
        errors.Add("Name is required.");
    if (string.IsNullOrWhiteSpace(request.Email) || !new EmailAddressAttribute().IsValid(request.Email))
        errors.Add("A valid email is required.");
    if (string.IsNullOrWhiteSpace(request.Message) || request.Message.Trim().Length < 10)
        errors.Add("Message should be at least 10 characters.");

    if (errors.Count > 0)
        return Results.BadRequest(new ContactResponse(false, string.Join(" ", errors)));

    store.Add(request with
    {
        Name = request.Name.Trim(),
        Email = request.Email.Trim(),
        Message = request.Message.Trim()
    });

    // In production, wire SMTP / queue here. For now we accept and acknowledge.
    _ = options.Value.Email;
    return Results.Ok(new ContactResponse(
        true,
        "Message received. I'll get back to you soon."));
});

app.MapFallbackToFile("index.html");

app.Run();
