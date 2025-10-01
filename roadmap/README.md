# Loqa Roadmap

## Vision

Loqa aims to become the leading open-source platform for local-first ambient intelligence, enabling privacy-preserving AI experiences that run entirely on user-controlled hardware.

## Milestones

### Phase 1: Foundation (Q3-Q4 2025)
**Status**: In Progress

- [x] Project inception and architecture design
- [x] Organization setup (Loqa Labs)
- [ ] Core runtime implementation
- [ ] Basic STT → LLM → TTS pipeline
- [ ] Initial clustering support
- [ ] Developer documentation

### Phase 2: Voice Interface (Q1 2026)
**Status**: Planned

- [ ] Optimized voice processing pipeline (<200ms latency)
- [ ] Wake word detection
- [ ] Multi-language support
- [ ] Basic skill system
- [ ] Home Assistant integration
- [ ] Alpha release for early adopters

### Phase 3: Extensibility (Q2 2026)
**Status**: Planned

- [ ] Plugin architecture (WASM-based)
- [ ] Skill marketplace foundation
- [ ] Visual configuration interface
- [ ] Advanced clustering (auto-discovery)
- [ ] Memory and context persistence
- [ ] Beta release

### Phase 4: Multimodal (Q3 2026)
**Status**: Planned

- [ ] Visual display support (e-ink, OLED)
- [ ] Touch input interfaces
- [ ] Gesture recognition
- [ ] Ambient notifications
- [ ] Mobile companion apps
- [ ] 1.0 stable release

### Phase 5: Intelligence (Q4 2026 - 2027)
**Status**: Future

- [ ] Local learning and adaptation
- [ ] Proactive assistance
- [ ] Multi-user personalization
- [ ] Advanced automation workflows
- [ ] Enterprise features
- [ ] Hardware partnerships

## Technical Priorities

### Near Term (Next 3 months)
1. **Performance**: Sub-200ms end-to-end latency
2. **Reliability**: 99.9% uptime for local services
3. **Modularity**: Clean service boundaries
4. **Developer Experience**: Simple SDK and clear docs

### Medium Term (3-6 months)
1. **Scalability**: Support 10+ node clusters
2. **Security**: End-to-end encryption for inter-node communication
3. **Extensibility**: Rich plugin ecosystem
4. **Compatibility**: Support for diverse hardware

### Long Term (6-12 months)
1. **Intelligence**: Context-aware responses
2. **Automation**: Complex workflow orchestration
3. **Federation**: Connect multiple Loqa instances
4. **Sustainability**: Open-core business model

## Community Involvement

### How to Contribute

- **Code**: Submit PRs for features and fixes
- **Testing**: Try alpha/beta releases and report issues
- **Skills**: Develop and share custom skills
- **Documentation**: Improve guides and examples
- **Hardware**: Test on diverse platforms
- **Feedback**: Share use cases and requirements

### Request for Comments (RFCs)

Major features go through the RFC process. Current RFCs:

- RFC-001: Plugin Architecture Design (Draft)
- RFC-002: Clustering Protocol Specification (Planned)
- RFC-003: Skill Definition Language (Planned)

## Success Metrics

### Technical
- Response latency < 200ms
- Memory usage < 500MB per service
- CPU usage < 20% idle
- 99.9% local availability

### Community
- 1000+ GitHub stars
- 100+ contributors
- 50+ community skills
- 10+ hardware platforms supported

### Impact
- 10,000+ active deployments
- 5+ enterprise customers
- 3+ hardware partnerships
- Privacy-first AI becoming mainstream

## Open-Core Model

### Forever Open Source
- Core runtime and control plane
- Basic STT, LLM, TTS services
- Standard skill library
- Community features

### Premium Features (Planned)
- Cloud backup and sync
- Advanced analytics
- Enterprise management
- Priority support
- Custom model training
- See [Workstream F value-add roadmap](workstream-f/value_add_roadmap.md) for the detailed plan.

## Get Involved

- Star the [repository](https://github.com/loqalabs/loqa-core)
- Join [discussions](https://github.com/loqalabs/loqa-meta/discussions)
- Follow [@loqalabs](https://github.com/loqalabs)
- Email: hello@ambiware.ai

---

*This roadmap is subject to change based on community feedback and technical discoveries. Last updated: September 2025*
