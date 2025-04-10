"use client";

import { useState } from "react";
import Image from "next/image";
import "./work.css";

interface Project {
  id: number;
  title: string;
  category: string;
  image: string;
}

/* ───── demo data ───── */
const base: Project[] = [
  { id: 1, title: "Sharlee",         category: "Branding",     image: "/projects/sharlee.jpg" },
  { id: 2, title: "Act Responsable", category: "Web Dev",      image: "/projects/act.jpg" },
  { id: 3, title: "Dua Lipa",        category: "Portrait",     image: "/projects/dua.jpg" },
  { id: 4, title: "Cocolyze",        category: "UX/UI Design", image: "/projects/cocolyze.jpg" },
  { id: 5, title: "Les Indécis",     category: "Branding",     image: "/projects/indecis.jpg" },
];

/* clone the demo list a few times to prove scrolling works */
const projects: Project[] = Array.from({ length: 10 }, () => base).flat();

export default function Work() {
  const [current, setCurrent] = useState<Project>(projects[0]);

  return (
    <div className="work-container">
      {/* ─────────── Left : preview ─────────── */}
      <section className="work-preview">
        <Image
          src={current.image}
          alt={current.title}
          fill
          priority
          className="work-preview-img"
        />
        <h2 className="work-preview-title">{current.title}</h2>
      </section>

      {/* ─────────── Right : list ─────────── */}
      <section className="work-list">
        <header className="work-list-header">
          <h2>Work</h2>
          <span>{projects.length}</span>
        </header>

        <div className="work-scroll">
          <ul className="work-items">
            {projects.map((p, idx) => (
              <li
                key={`${p.id}-${idx}`} /* guarantees uniqueness */
                className={`work-item ${
                  p === current ? "work-item-active" : ""
                }`}
                onMouseEnter={() => setCurrent(p)}
                onClick={() => setCurrent(p)}
              >
                <span className="work-item-title">{p.title}</span>
                <span className="work-item-cat">{p.category}</span>
              </li>
            ))}
          </ul>
        </div>
      </section>
    </div>
  );
}
