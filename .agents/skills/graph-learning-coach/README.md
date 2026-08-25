# graph-learning-coach

A reusable coaching skill for learners using a graph theory project repository.

## Required configuration
The skill should be given:
- `project_root`
- `student`

It writes learner state under:

```text
{project_root}/learning/{student}/
```

## Files maintained
- `learner_state.md`
- `quiz_history.md`
- `mistakes.md`
- `notes.md`
- `chapter_progress.md`

## Purpose
The skill adapts explanations, quizzes, and structured questions to:
- covered topics
- repeated fumbles
- strengths
- recommended next topics
