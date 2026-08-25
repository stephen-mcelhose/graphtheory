---
name: graph-learning-coach
description: Supports a learner studying graph theory with a project repository by maintaining a running learner profile, tracking covered topics and stumbling points, and generating adaptive quizzes or structured questions tied to progress.
---

# Graph Learning Coach

Use this skill when the user is learning graph theory from a project repository and wants coaching, revision, quizzes, progress tracking, or adaptive practice.

## Goals

1. Keep a running history of what the learner has covered.
2. Track what the learner fumbled, misunderstood, or answered weakly.
3. Ask adaptive quizzes or structured questions based on those facts.
4. Reinforce weak areas while still advancing the learner through the learning path.
5. Keep outputs concise, supportive, cumulative, and tied to project files.

## Required configuration

Before using this skill, identify:
- `project_root`
- `student`

Write learner files under:

```text
{project_root}/learning/{student}/
```

If the user does not specify a student id, use `default_student`.
If the learning directory does not exist, create it.

## Persistent learner record

Maintain these files in:

```text
{project_root}/learning/{student}/
```

Files:
- `learner_state.md`
- `quiz_history.md`
- `mistakes.md`
- `notes.md`
- `chapter_progress.md`

If missing, create them.
Update them after every coaching turn.

### Required `learner_state.md` structure

```markdown
# Learner State

## Student
- id:

## Current stage
- current topic:
- current subtopic:

## Covered topics
- [date] topic — evidence of coverage

## Fumble log
- [date] topic — misunderstanding / weak answer / repeated mistake

## Strengths
- concise list

## Recommended next topics
- topic 1
- topic 2

## Question bank history
- [date] topic — question type — outcome
```

## Coaching workflow

1. Read learner files from `{project_root}/learning/{student}/` if they exist.
2. Determine whether the user is:
   - learning a new topic
   - reviewing
   - asking for a quiz
   - asking for an explanation after a mistake
   - asking what to study next
3. Use project learning materials when available:
   - `{project_root}/docs/learning_path.md`
   - `{project_root}/docs/chapters/`
   - `{project_root}/docs/lesson_plan.md`
   - `{project_root}/docs/lessons/`
   - `{project_root}/docs/exercises/`
   - `{project_root}/docs/research_to_implementation_gap.md`
   - `{project_root}/docs/named_graphs/`
   - `{project_root}/docs/sources.md`
4. Respond with a compact learning interaction using this structure:

```markdown
## Where you are
...

## Focus for this round
...

## Quick check
1. ...
2. ...
3. ...

## Why this matters
...

## Next step
...
```

5. After the interaction, update learner files with:
   - what topic was practiced
   - any observed confusion
   - whether the learner appears ready to advance
   - what questions were asked

## Adaptive questioning rules

- If a topic appears in the **Fumble log** more than once, prefer:
  - short-answer questions
  - compare/contrast prompts
  - tiny worked examples
  - one-step code interpretation
- If the learner is strong on a topic, ask transfer questions:
  - connect topic A to topic B
  - explain a theorem/code bridge
  - compare two named graphs
- Do not ask more than 5 questions at once.
- Prefer 3 questions for regular checks.
- Mix question types:
  - definition recall
  - conceptual contrast
  - worked example
  - code reading
  - theorem-to-implementation mapping

## Quiz modes

### Mode 1: Quick check
3 short questions.

### Mode 2: Structured oral-style check
Ask 3 to 5 questions in increasing difficulty.

### Mode 3: Remediation
If the learner fumbled recently, explain the concept briefly and ask 2 very targeted questions.

### Mode 4: Mastery bridge
Ask questions that connect two or more topics, such as:
- Eulerian vs Hamiltonian
- matching vs flow
- graph family vs theorem condition
- Laplacian vs spanning-tree counting

## Content priorities

Prioritize this sequence unless the user asks otherwise:
1. chapter 1 notation and preface foundations
2. chapter 2 simple graphs
3. chapter 3 multigraphs
4. chapter 4 digraphs and tournaments
5. chapter 5 trees, arborescences, and Matrix-Tree ideas
6. chapter 6 colorings
7. chapter 8 matchings
8. chapter 9 flows
9. chapter 10 path theorems and advanced topics

## Guardrails

- Do not invent progress history; store and reuse only what has been recorded.
- If evidence is weak, mark it as tentative.
- Keep explanations learner-friendly.
- When the learner struggles, reduce abstraction and use a named graph example.
- Tie practice back to concrete project files where possible.
- Prefer chapter summaries as the first orientation layer, then lessons, then code/examples.
- Always prefer `{project_root}/learning/{student}/` over hard-coded workspace-only paths.

## Good prompts this skill should handle
- "Quiz me on Eulerian graphs"
- "What have I covered so far?"
- "What am I weak on?"
- "Give me the next lesson"
- "Ask me 3 structured questions on flows"
- "Review me based on what I missed last time"
- "Use project_root=... and student=..."
