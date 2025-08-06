# Stakeholder Lenses for Viewing Microservices

![main](./images/LencesOfView/image.png)

This diagram illustrates that a microservices architecture is not viewed uniformly across an organization. Instead, different roles interact with it through unique "lenses," each with its own set of priorities and concerns.

## Leadership & Strategic View
- `Non-Technical Founder / Engineer Director`: Focuses on the organizational impact. Their primary concerns are team structure, division of responsibilities, managing team growth (e.g., offshore teams), and setting Key Performance Indicators (KPIs).

- `Technical Founder (CTO)`: Takes a high-level view of the system's technical health and business impact. They prioritize uptime (availability), feature velocity, scalability, reliability, and security.

## Development & Implementation View
- `Indie Developer`: Focuses on the initial project architecture and workflow. Their main task is to divide work into services and figure out how to distribute that work across developers.

- `Software Engineer`: Has a focused, hands-on perspective. They are concerned with contributing to a small section of the software, the specific lifecycle of that small section, and how to access data from other microservices and their databases.

## Architecture & Operations View
- `Tech Architect`: Responsible for the system's foundation. They decide what technology to build with, how to set up communications between services, how to ensure long-term maintainability, and what cloud infrastructure to use.

- `DevOps`: Focuses on the operational reality of running the system. As shown in the diagram, their concerns include how changes and deployments will work and creating failure scenario mapping to ensure resilience.