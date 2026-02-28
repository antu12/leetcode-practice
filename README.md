## Summary

Each problem folder contains the solution implementation, unit tests, and the problem statement.

```
Problems/
├── problem1/
│   ├── problem1.go
│   ├── problem1_test.go
│   └── question.md
```

### Creating New Problems

Use the `new_problem.sh` script to scaffold a new problem:

```bash
./new_problem.sh <problem_name>
```

Example:
```bash
./new_problem.sh two_sum
```

This creates a new problem directory with:
- `<problem_name>.go` - Solution template
- `<problem_name>_test.go` - Unit test template
- `question.md` - Problem statement file