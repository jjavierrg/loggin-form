# ---------- FINAL (SCRATCH) ----------
FROM scratch

# Copiar binario y certificados (construidos en la pipeline de CI/CD)
COPY loginapp /loginapp
COPY server.crt /server.crt
COPY server.key /server.key

EXPOSE 80 443

# Healthcheck → usa el binario para testear http://localhost/health
HEALTHCHECK --interval=10s --timeout=2s --retries=3 CMD ["/loginapp", "health"]

ENTRYPOINT ["/loginapp"]
