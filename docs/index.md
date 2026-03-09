---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: "UltraHooks"
  text: "The Universal Git Hooks Manager"
  tagline: Blazingly fast, lightweight, and language-agnostic Git hooks management in Go.
  actions:
    - theme: brand
      text: Get Started
      link: /README
    - theme: alt
      text: Installation
      link: /installation

features:
  - title: Blazingly Fast
    details: Instant execution times. Zero bloated runtime dependencies. Supports concurrent parallel hook execution via goroutines.
  - title: Language Agnostic
    details: Works universally for Go, Node.js, Python, Rust, PHP, C++, and more. Automatically discovers .sh, .bat, and .ps1 proxies natively.
  - title: Extremely Clean
    details: Never pollutes your root .git/hooks directory again. Surgical precision installation directly wires hooks from your shared repository configurations.
---

