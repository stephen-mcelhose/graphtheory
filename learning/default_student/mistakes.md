# Mistakes

- [2026-08-26] SCC grouping — when asked to find SCCs of `A→B, B→A, B→C`, listed "A, B and C" as three separate components. Correct answer: {A,B} is one SCC (mutually reachable), {C} is a second SCC (cannot return to B). Key reminder: SCC groups vertices by mutual reachability, not individual listing.
