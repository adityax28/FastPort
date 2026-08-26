using Aditya.Portfolio.Models;
using Microsoft.Extensions.Options;

namespace Aditya.Portfolio.Services;

public sealed class ContactStore
{
    private readonly List<(DateTimeOffset At, ContactRequest Request)> _messages = new();
    private readonly object _lock = new();

    public void Add(ContactRequest request)
    {
        lock (_lock)
        {
            _messages.Add((DateTimeOffset.UtcNow, request));
        }
    }

    public int Count
    {
        get
        {
            lock (_lock)
            {
                return _messages.Count;
            }
        }
    }
}

public sealed class PortfolioData
{
    private readonly PortfolioOptions _options;
    private readonly DateTimeOffset _startedAt = DateTimeOffset.UtcNow;

    public PortfolioData(IOptions<PortfolioOptions> options)
    {
        _options = options.Value;
    }

    public DateTimeOffset StartedAt => _startedAt;

    public ProfilePayload GetProfile() => new(
        Name: _options.Name,
        Title: _options.Title,
        Tagline: _options.Tagline,
        YearsExperience: _options.YearsExperience,
        Domain: _options.Domain,
        Email: _options.Email,
        GitHub: _options.GitHub,
        LinkedIn: _options.LinkedIn,
        Focus:
        [
            "Transaction & payment systems",
            "API integrations at production scale",
            "Database & performance optimization",
            "Owning bottlenecks end-to-end"
        ],
        Projects:
        [
            new ProjectItem(
                "pay-core",
                "Payment transaction pipeline",
                "Backend flows for payment initiation, status tracking, and reconciliation-oriented processing in a fintech environment.",
                ["C#", ".NET", "SQL", "APIs"],
                "Reliable handling of high-value money movement paths."),
            new ProjectItem(
                "api-mesh",
                "Upstream / downstream integrations",
                "Hardened API integrations with external and internal services—timeouts, retries, and clear failure surfaces for operators.",
                [".NET", "HTTP", "CI/CD"],
                "Fewer silent failures; faster incident diagnosis."),
            new ProjectItem(
                "db-perf",
                "Query & throughput optimization",
                "Identified hot paths, tightened queries and indexes, and reduced latency under production load.",
                ["SQL Server", "Profiling", "Caching"],
                "Measurable wins on critical transaction endpoints."),
            new ProjectItem(
                "ops-ownership",
                "Production problem solving",
                "Took ownership beyond tickets: mapped business impact, isolated bottlenecks, and shipped durable fixes.",
                ["Observability", "Incident response"],
                "From firefighting to scalable systems thinking.")
        ],
        Stack:
        [
            "C#", ".NET", "ASP.NET Core", "SQL Server", "REST APIs",
            "CI/CD", "Git", "Caching", "Logging", "Fintech domain"
        ]);

    public SystemStatus GetStatus(IHostEnvironment env) => new(
        Status: "operational",
        Environment: env.EnvironmentName,
        StartedAt: _startedAt,
        UptimeSeconds: (DateTimeOffset.UtcNow - _startedAt).TotalSeconds,
        Runtime: System.Runtime.InteropServices.RuntimeInformation.FrameworkDescription,
        Pipelines: ["build", "test", "deploy"]);
}
