# Build
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src

COPY Aditya.Portfolio/Aditya.Portfolio.csproj Aditya.Portfolio/
RUN dotnet restore Aditya.Portfolio/Aditya.Portfolio.csproj

COPY Aditya.Portfolio/ Aditya.Portfolio/
WORKDIR /src/Aditya.Portfolio
RUN dotnet publish -c Release -o /app/publish --no-restore

# Runtime
FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS final
WORKDIR /app

ENV ASPNETCORE_ENVIRONMENT=Production
ENV ASPNETCORE_URLS=http://0.0.0.0:8080

COPY --from=build /app/publish .

EXPOSE 8080

# Render sets $PORT; fall back to 8080 locally
CMD ["sh", "-c", "ASPNETCORE_URLS=http://0.0.0.0:${PORT:-8080} exec dotnet Aditya.Portfolio.dll"]
