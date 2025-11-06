# Overview

As a software engineer focused on expanding my expertise in modern communication systems, I developed a comprehensive Messaging API that enables businesses to manage customer communications effectively through multiple channels.

## Project Description

The Messaging API is a robust platform that allows users to select subscription plans and define their preferred communication methods with customers, whether through email or SMS. The system integrates with industry-leading services including Twilio (for SMS), SendGrid (for email), and Ntfy (for notifications) to handle message delivery. A key feature is its comprehensive logging system that tracks every message with detailed status updates (open, read, delivered, sent, error, etc.).

The main purpose of developing this software was to create a unified communication platform that:
1. Simplifies multi-channel customer communication
2. Provides reliable message delivery tracking
3. Offers flexible integration options
4. Maintains detailed message status history
5. Supports scalable subscription-based access

[Software Demo Video](https://youtu.be/cOuzPUX_40o)

# Development Environment

## Tools Used
- **Go**: Primary programming language
- **Docker**: Container management for local PostgreSQL database
- **Git**: Version control system
- **Ngrok**: Secure tunnel for webhook testing
- **API Integrations**:
  - Twilio: SMS messaging
  - SendGrid: Email delivery
  - Ntfy: Push notifications

The API is built using Go with a carefully structured architecture:

- **Web Framework**: Gin for handling HTTP requests
- **Database**: PostgreSQL with GORM ORM for data persistence
- **Architecture**: Clean architecture pattern with:
  - Entities: Core business objects
  - Repositories: Data access layer
  - Use Cases: Business logic
  - Handlers: Request processing
  - Factories: Integration management

This architecture ensures clean separation of concerns and makes the system extensible for future integrations.

# Useful Websites

- [Twilio Documentation](https://www.twilio.com/docs) - Comprehensive guides for SMS integration
- [SendGrid API Reference](https://sendgrid.com/en-us) - Email delivery and webhook documentation
- [Ngrok Documentation](https://ngrok.com/) - Webhook testing and tunnel setup
- [Ntfy Documentation](https://ntfy.sh/) - Push notification implementation
- [Go Gin Framework](https://gin-gonic.com/) - Web framework documentation
- [GORM Documentation](https://gorm.io/) - Database operations and modeling

# Future Work

- **Bulk Messaging**: Support for sending messages to multiple recipients
- **Message Scheduling**: Add capability to schedule messages for future delivery
- **Webhook Retries**: Implement retry mechanism for failed webhook deliveries
