# Projects

## SSH Portfolio
*This very application you're using!*

A unique SSH-based portfolio built ground-up in Go with Bubble Tea. BUT, this was [AI assisted](https://antigravity.google/) (I PROMISE I'll do better next time). I just wanted to bring my vision of my portfolio to life so I can showcase the things I did actually build myself (see below).

**Features:**
- Password authentication
- Rate limiting
- Beautiful TUI with vim keybindings
- Markdown content rendering
- Docker deployment

**Tech Stack:**
- Go 1.26.0
- Bubble Tea, Lip Gloss, Bubbles (yes, these are real libraries)
- SQLite for persistence
- GitHub Actions CI/CD

[GitHub Repo](https://github.com/adamdeleeuw/ssh-portfolio)

---

## Multi-threaded HTTP Server

**Description**: Currently building this. Will hopefully be done soon (maybe? Haven't made much progress).

[GitHub Repo](https://github.com/adamdeleeuw/cpp-multithreaded-server)

---

## Heap Memory Allocator - UBC (CPEN 212)

**Description:** This may be shocking, but it allocates memory to the heap ;)

**Tech:** C, and that's about it. Edited the provided Makefiles and wrote a bash script to automate the Lua test traces.

**Highlights:**
- alloc() & free() - Implemented fragmentation-reducing algorithms including immediate bidirectional coalescing and block splitting. 
- realloc() - Engineered in-place reallocation to minimize expensive memcpy operations.
- Integrity - Developed a comprehensive heap consistency checker to validate invariants.

No repo link because I think my [prof](https://ece.ubc.ca/alexandra-sasha-fedorova/) and the Dean would smack me (I'm kidding in case they see this...).

---

## Tron Light Cycle Game - UBC (CPEN 211)

**Description:** A real-time game engine built on a Nios V processor. It features zero-latency responsiveness by replacing inefficient polling with an interrupt-driven architecture.

**Tech:** C, RISC-V, and FPGA (DE10-Lite).

**Highlights:**
- Hardware Integration: Managed Direct Memory Access (DMA) to VGA buffers and memory-mapped I/O (LEDs/Hex) using optimized bitwise operations. 
- Autonomous Agent: Developed a predictive AI that analyzes pixel data for collision avoidance. 
- State Management: Implemented complex state machines for win conditions and real-time scoring.

Again, no github repo ([Dr. Lemieux](https://ece.ubc.ca/guy-lemieux/) said it was "too good" to be on github, or maybe he just didn't want to see my name on it lol).

---

## Buffers, Concurrency, and Wikipedia Tool - UBC (CPEN 221, AI permitted)

**Description:** A Java project that combines a thread-safe, time-expiring cache with a Wikipedia mediator service. It parses client requests, validates parameters, fetches and caches page content, and reports usage statistics through a secure client/server interface.

**Tech:** Java, JUnit 5, ANTLR, Gson.

**Highlights:**

- Caching: Implements a fixed-size, time-expiring buffer that supports concurrent access, eviction, and freshness checks for cached entries.
- Request Analytics: Tracks request frequency and timing to provide zeitgeist (most requested) and peak load (max requests in a window) metrics.
- Parsing & Protocol: Uses a JSON-based command protocol with encryption, robust parsing, and error handling for client requests.

No public repo for similar reasons as above.

---

## Java Graphs & Applications - UBC (CPEN 221, AI permitted)

**Description:** A comprehensive library for graph data structures and algorithms with applications in text similarity analysis and terrain analysis. Features multiple graph implementations, advanced graph algorithms, and practical applications including document similarity computation.

**Tech:** Java, JUnit 5, and Git.

**Highlights:**

- Text Similarity Analysis: Implemented document processing pipeline with word frequency analysis and Jensen-Shannon Divergence (JSD) computation for comparing textual documents from files and URLs.
- Graph Algorithms: Developed core graph algorithms including graph partitioning, shortest path computation, and minimum spanning tree operations on both adjacency list and adjacency matrix representations.

No public repo.

---

## Java Image Processing & Analysis Tool - UBC (CPEN 221, AI permitted)

**Description:** A robust tool for digital signal processing and image manipulation. It includes a document alignment engine and background replacement functionality. 

**Tech:** Java (JDK 17), JUnit 5, and Git. 

**Highlights:**

- Signal Processing: Optimized document alignment using 2D Discrete Fourier Transforms (DFT) and Frequency Domain analysis, reducing time complexity via dynamic down-sampling. 
- Algorithmic Safety: Engineered a Green Screen tool using Stack-based Iterative DFS to identify connected components, explicitly avoiding recursive StackOverflow Errors. 
- Statistical Matching: Implemented Vector Space Models (Cosine Similarity) for image matching and enforced reliability with comprehensive unit testing.

No public repo.

---

## UBC Finds (Contributor) - Campus Utility Tracker (lowkey vibe coded, I don't claim to know TS or React)

**Description:** A real-time, community-driven utility tracker for the University of British Columbia. It helps students locate and report the status of campus utilities including water fountains, bike racks, food vendors, bus stops, and emergency phones across UBC's campus.

**Tech:** Next.js, React, TypeScript, Supabase, Google Maps API, and Tailwind CSS.

**Highlights:**

- Interactive Mapping: Integrated Google Maps API to display 300+ campus utilities with real-time status updates and location coordinates, enabling students to quickly find nearby amenities.
- Community Reporting System: Built a crowd-sourced issue reporting system using Supabase backend, allowing students to submit and view utility status reports with automatic yellow marker indicators for flagged items.
- Smart Filtering Interface: Developed category-based filtering (water, bike, food, bus, emergency) with instant map updates, helping users focus on specific utility types through an intuitive toggle system.
- Responsive Onboarding: Created a step-by-step desktop onboarding modal with filter controls, search functionality, and reporting features to guide new users through the platform's capabilities.

[GitHub Repo](https://github.com/UBCFinds/ubcfinds)

---

## Learning & Exploration Projects

Projects I built following tutorials or for learning purposes - because learning in public is important.

### Flappy Bird AI with NEAT (Tutorial Project)

**Description:** An intelligent Flappy Bird implementation using neuroevolution through the NEAT (NeuroEvolution of Augmenting Topologies) algorithm. Followed [this tutorial](https://www.youtube.com/watch?v=MMxFDaIOHsE) to learn about evolutionary algorithms and game AI.

**Tech:** Python, Pygame, NEAT-Python, Matplotlib, Poetry.

**What I Learned:**
- Evolutionary algorithms and neuroevolution
- How fitness functions drive emergent behaviors
- Python's Pygame library for game development

[GitHub Repo](https://github.com/adamdeleeuw/flappy-bird-ai)

---

*More projects coming soon...*
