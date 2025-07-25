# Как код, выглядящий просто, может оказаться сложным

## Без конфигурационного файла (хардкод)

```csharp
using System;

public class NotificationService
{
    public void SendNotification(string message)
    {
        // Жёстко заданные настройки
        string smtpServer = "smtp.example.com";
        int port = 587;
        string username = "user@example.com";
        string password = "secret";
        bool enableSsl = true;
        int timeout = 30;
        string senderEmail = "noreply@example.com";
        
        // Логика отправки
        Console.WriteLine($"Sending notification via {smtpServer}:{port}");
        Console.WriteLine($"SSL: {enableSsl}, Timeout: {timeout}s");
        Console.WriteLine($"From: {senderEmail}");
        Console.WriteLine($"Message: {message}");
        
        // Реальная отправка email...
    }
}

class Program
{
    static void Main()
    {
        var service = new NotificationService();
        service.SendNotification("Hello, this is a test notification!");
    }
}
```

## Решение с конфигурацией

```csharp
using System;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Options;

// 1. Конфигурационный класс
public class SmtpConfig
{
    public string Server { get; set; } = "smtp.default.com";
    public int Port { get; set; } = 25;
    public string Username { get; set; } = string.Empty;
    public string Password { get; set; } = string.Empty;
    public bool EnableSsl { get; set; } = true;
    public int Timeout { get; set; } = 30;
    public string SenderEmail { get; set; } = "noreply@default.com";
}

// 2. Сервис с зависимостями
public class NotificationService
{
    private readonly SmtpConfig _config;

    public NotificationService(IOptions<SmtpConfig> config)
    {
        _config = config.Value;
        ValidateConfig();
    }

    public void SendNotification(string message)
    {
        Console.WriteLine($"Sending notification via {_config.Server}:{_config.Port}");
        Console.WriteLine($"SSL: {_config.EnableSsl}, Timeout: {_config.Timeout}s");
        Console.WriteLine($"From: {_config.SenderEmail}");
        Console.WriteLine($"Message: {message}");
        
        // Реальная отправка email...
    }

    private void ValidateConfig()
    {
        if (string.IsNullOrWhiteSpace(_config.Server))
            throw new ArgumentNullException(nameof(SmtpConfig.Server));
        
        if (_config.Port < 1 || _config.Port > 65535)
            throw new ArgumentOutOfRangeException(nameof(SmtpConfig.Port));
        
        if (string.IsNullOrWhiteSpace(_config.SenderEmail))
            throw new ArgumentNullException(nameof(SmtpConfig.SenderEmail));
    }
}

class Program
{
    static void Main()
    {
        // 3. Настройка конфигурации
        var config = new ConfigurationBuilder()
            .AddJsonFile("appsettings.json")
            .Build();
        
        // 4. Настройка DI
        var services = new ServiceCollection();
        
        services.AddOptions<SmtpConfig>()
            .Bind(config.GetSection("Smtp"))
            .Validate(c => !string.IsNullOrWhiteSpace(c.Server), "Server is required")
            .Validate(c => c.Port > 0 && c.Port <= 65535, "Invalid port number")
            .Validate(c => !string.IsNullOrWhiteSpace(c.SenderEmail), "Sender email is required");
        
        services.AddSingleton<NotificationService>();
        
        var provider = services.BuildServiceProvider();
        
        // 5. Использование сервиса с обработкой ошибок
        try
        {
            var service = provider.GetRequiredService<NotificationService>();
            service.SendNotification("Hello from configured service!");
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error: {ex.Message}");
        }
    }
}
```

appsettings.json:
```json
{
  "Smtp": {
    "Server": "smtp.company.com",
    "Port": 587,
    "Username": "service@company.com",
    "Password": "secure_password",
    "EnableSsl": true,
    "Timeout": 20,
    "SenderEmail": "notifications@company.com"
  }
}
```