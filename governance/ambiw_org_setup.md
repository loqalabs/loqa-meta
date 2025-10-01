# Loqa Labs Setup Notes

## Organization
- **GitHub handle:** `loqalabs`
- **Display name:** Loqa Labs
- **Description:** "Building Loqa: a local-first, open-core ambient intelligence platform."
- **Website:** https://ambiware.ai
- **Contact email:** hello@ambiware.ai
- **Brand pairing:** Loqa remains the flagship open-source project; Loqa Labs is the parent org.

## DNS Configuration (Porkbun)

### GitHub Pages
- Remove default Pixie records (`ALIAS @ → pixie.porkbun.com` and wildcard CNAME) before adding the following.
- Add A records for `@` pointing to GitHub Pages:
  - `185.199.108.153`
  - `185.199.109.153`
  - `185.199.110.153`
  - `185.199.111.153`
- Add `CNAME www → loqalabs.github.io`.
- In the Pages repo: add a `CNAME` file with `ambiware.ai`, set the custom domain in GitHub Pages, and enable HTTPS once available.

### Email (iCloud Custom Domain)
- MX records (host `@`):
  - `10 mx01.mail.icloud.com.`
  - `10 mx02.mail.icloud.com.`
- SPF TXT (host `@`):
  - `"v=spf1 include:icloud.com ~all"`
- DKIM: add the `sig1._domainkey` CNAME provided in the iCloud setup panel (value unique per account).
- Create aliases such as `hello@ambiware.ai` for public contact; optionally forward to personal inbox.

## Repository Strategy
- **Core monorepo:** `loqa-core` (runtime, control plane, SDKs, docs, proto schemas).
- **Satellite repos:**
  - `loqa-skills` – community skill templates and curated examples.
  - `loqa-site` – marketing/landing site served via GitHub Pages.
  - `loqa-meta` – governance, RFCs, org-wide documentation.
  - (Future) `ambiware-pro` – private/premium extensions.
- Configure shared workflows, CODEOWNERS, branch protections, and issue templates after scaffolding.

## Immediate Next Steps
1. Update org profile with avatar, bio, and website link.
2. Scaffold repositories via `gh repo create ...` commands; seed each with README, LICENSE, CONTRIBUTING, CODE_OF_CONDUCT.
3. Stand up initial GitHub Pages repo under `loqa-site` and connect to `ambiware.ai` once DNS propagates.
4. Document governance and roadmap in `loqa-meta`, announcing the Loqa Labs ↔ Loqa relationship.
