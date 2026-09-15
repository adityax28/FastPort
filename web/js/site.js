const TOKENS = [
  "C#", ".NET", "CI/CD", "SQL", "Postgres", "Redis",
  "K8s", "Nginx", "gRPC", "APIs", "JVM", "Docker",
  "Payments", "ASP.NET", "Git", "Observability"
];

function formatUptime(seconds) {
  const s = Math.max(0, Math.floor(seconds));
  const h = Math.floor(s / 3600);
  const m = Math.floor((s % 3600) / 60);
  const r = s % 60;
  if (h > 0) return `${h}h ${m}m`;
  if (m > 0) return `${m}m ${r}s`;
  return `${r}s`;
}

async function fetchJson(url, options) {
  const res = await fetch(url, options);
  const data = await res.json().catch(() => ({}));
  if (!res.ok) {
    const msg = data.message || `Request failed (${res.status})`;
    throw new Error(msg);
  }
  return data;
}

function renderFocus(items) {
  const list = document.getElementById("focus-list");
  list.innerHTML = items
    .map(
      (text, i) => `
      <li>
        <span class="idx">${String(i + 1).padStart(2, "0")}</span>
        <span>${text}</span>
      </li>`
    )
    .join("");
}

function renderWork(projects) {
  const root = document.getElementById("work-stream");
  root.innerHTML = projects
    .map(
      (p) => `
    <article class="work-item">
      <p class="work-id">${p.id}</p>
      <h3>${p.title}</h3>
      <p>${p.summary}</p>
      <p class="work-impact">${p.impact}</p>
      <div class="work-tags">${p.tags.map((t) => `<span>${t}</span>`).join("")}</div>
    </article>`
    )
    .join("");
}

function renderStack(stack) {
  const rail = document.getElementById("stack-rail");
  rail.innerHTML = stack.map((s) => `<span>${s}</span>`).join("");
}

function applyProfile(profile) {
  document.getElementById("role-line").textContent =
    `${profile.title} · ${profile.domain}`;
  document.getElementById("tagline").textContent = profile.tagline;
  document.getElementById("years").textContent = profile.yearsExperience;

  const email = profile.email;
  const mail = document.getElementById("link-email");
  mail.textContent = email;
  mail.href = `mailto:${email}`;

  document.getElementById("link-github").href = profile.gitHub || "#";
  document.getElementById("link-linkedin").href = profile.linkedIn || "#";

  renderFocus(profile.focus);
  renderWork(profile.projects);
  renderStack(profile.stack);
}

function applyStatus(status) {
  const pill = document.getElementById("status-pill");
  const label = document.getElementById("status-label");
  pill.classList.add("live");
  label.textContent = status.status;

  document.getElementById("env-label").textContent =
    `env://${(status.environment || "local").toLowerCase()}`;
  document.getElementById("uptime-label").textContent =
    `uptime ${formatUptime(status.uptimeSeconds)}`;
  document.getElementById("runtime-label").textContent =
    status.runtime || "runtime —";
}

function startUptimeTicker(startedAtIso) {
  const started = new Date(startedAtIso).getTime();
  const tick = () => {
    const seconds = (Date.now() - started) / 1000;
    document.getElementById("uptime-label").textContent =
      `uptime ${formatUptime(seconds)}`;
  };
  tick();
  setInterval(tick, 1000);
}

function initTokenField() {
  const canvas = document.getElementById("token-field");
  const ctx = canvas.getContext("2d");
  let width = 0;
  let height = 0;
  let dpr = 1;
  let particles = [];
  let raf = 0;

  function resize() {
    dpr = Math.min(window.devicePixelRatio || 1, 2);
    width = window.innerWidth;
    height = window.innerHeight;
    canvas.width = Math.floor(width * dpr);
    canvas.height = Math.floor(height * dpr);
    canvas.style.width = `${width}px`;
    canvas.style.height = `${height}px`;
    ctx.setTransform(dpr, 0, 0, dpr, 0, 0);
  }

  function spawn() {
    particles = TOKENS.map((text, i) => ({
      text,
      x: Math.random() * width,
      y: Math.random() * height,
      vx: (Math.random() - 0.5) * 0.18,
      vy: (Math.random() - 0.5) * 0.12 - 0.04,
      opacity: 0.08 + (i % 5) * 0.03,
      size: 11 + (i % 4) * 2
    }));
  }

  function frame() {
    ctx.clearRect(0, 0, width, height);
    ctx.font = "500 13px 'IBM Plex Mono', monospace";

    for (const p of particles) {
      p.x += p.vx;
      p.y += p.vy;
      if (p.x < -80) p.x = width + 40;
      if (p.x > width + 80) p.x = -40;
      if (p.y < -40) p.y = height + 20;
      if (p.y > height + 40) p.y = -20;

      ctx.globalAlpha = p.opacity;
      ctx.fillStyle = "#9eb6c8";
      ctx.font = `500 ${p.size}px 'IBM Plex Mono', monospace`;
      ctx.fillText(p.text, p.x, p.y);
    }

    ctx.globalAlpha = 1;
    raf = requestAnimationFrame(frame);
  }

  resize();
  spawn();
  frame();

  window.addEventListener("resize", () => {
    resize();
    spawn();
  });

  document.addEventListener("visibilitychange", () => {
    if (document.hidden) cancelAnimationFrame(raf);
    else raf = requestAnimationFrame(frame);
  });
}

function wireContactForm() {
  const form = document.getElementById("contact-form");
  const status = document.getElementById("form-status");
  const submit = document.getElementById("contact-submit");

  form.addEventListener("submit", async (e) => {
    e.preventDefault();
    status.className = "form-status mono";
    status.textContent = "transmitting…";
    submit.disabled = true;

    const payload = {
      name: form.name.value,
      email: form.email.value,
      message: form.message.value
    };

    try {
      const res = await fetchJson("/api/contact", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      });
      status.classList.add("ok");
      status.textContent = res.message || "sent";
      form.reset();
    } catch (err) {
      status.classList.add("err");
      status.textContent = err.message || "failed";
    } finally {
      submit.disabled = false;
    }
  });
}

async function boot() {
  document.getElementById("year").textContent = String(new Date().getFullYear());
  initTokenField();
  wireContactForm();

  try {
    const [profile, status] = await Promise.all([
      fetchJson("/api/profile"),
      fetchJson("/api/status")
    ]);
    applyProfile(profile);
    applyStatus(status);
    startUptimeTicker(status.startedAt);
  } catch {
    document.getElementById("status-label").textContent = "offline";
  }
}

boot();
