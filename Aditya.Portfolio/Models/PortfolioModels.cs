namespace Aditya.Portfolio.Models;

public sealed class PortfolioOptions
{
    public const string SectionName = "Portfolio";

    public string Name { get; set; } = "Aditya";
    public string Title { get; set; } = "Backend Developer";
    public string Tagline { get; set; } = string.Empty;
    public int YearsExperience { get; set; } = 2;
    public string Email { get; set; } = string.Empty;
    public string GitHub { get; set; } = string.Empty;
    public string LinkedIn { get; set; } = string.Empty;
    public string Domain { get; set; } = "Fintech";
}

public sealed record ContactRequest(
    string Name,
    string Email,
    string Message);

public sealed record ContactResponse(
    bool Ok,
    string Message);

public sealed record SystemStatus(
    string Status,
    string Environment,
    DateTimeOffset StartedAt,
    double UptimeSeconds,
    string Runtime,
    IReadOnlyList<string> Pipelines);

public sealed record ProjectItem(
    string Id,
    string Title,
    string Summary,
    string[] Tags,
    string Impact);

public sealed record ProfilePayload(
    string Name,
    string Title,
    string Tagline,
    int YearsExperience,
    string Domain,
    string Email,
    string GitHub,
    string LinkedIn,
    IReadOnlyList<string> Focus,
    IReadOnlyList<ProjectItem> Projects,
    IReadOnlyList<string> Stack);
