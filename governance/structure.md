# Governance Structure

```mermaid
graph TD
    subgraph Community
        Contributors --> RFC["RFC Draft"]
        Contributors --> PRs["Pull Requests"]
    end

    RFC --> Review["Maintainer / Community Review"]
    Review --> Decision{Decision}
    Decision -->|Accepted| Implementation["Implementation & Tracking"]
    Decision -->|Needs Rework| RFC
    Decision -->|Rejected| Archive["Archive / Notes"]

    Implementation --> Release["Release & Docs"]

    subgraph Loqa Studio Governance
        MarketplaceSubmit["Studio Submission"] --> Vetting["Safety & Quality Vetting"]
        Vetting -->|Approved| Catalog["Marketplace Catalog"]
        Vetting -->|Changes Requested| MarketplaceSubmit
    end

    Maintainers["Core Maintainers"] --> Review
    Maintainers --> Vetting
    FutureSteering["Future Steering Group"] -.-> Decision
```

## Roles & Responsibilities
- **Contributors** – draft RFCs, submit code/docs, share Studio add-ons.
- **Maintainers** – steward the roadmap, review RFCs, approve marketplace entries during the bootstrap phase.
- **Future Steering Group** – will be established to share decision-making once the community grows; today it’s represented by maintainers.

## Processes
- **RFC Flow:** Submission → Discussion/Review → Decision → Implementation → Release.
- **Marketplace Flow:** Submission → Vetting (quality, security, ethics) → Catalog publication.

See also:
- [Governance overview](README.md)
- [Licensing philosophy](licensing.md)
- [Privacy principles](privacy.md)
- [Marketplace guidelines](../studio/README.md)
