# Exercise 1: Summarization Pattern

------------------------------------------------------------------------

## Objective

Create a reusable prompt template for text summarization that works
consistently across different text types.

------------------------------------------------------------------------

## Weak Prompt (Original)

> "Summarize this text."

**Problem:** Too vague - no constraints on length, focus, or format.

------------------------------------------------------------------------

## Improved Pattern (Reusable Template)

> "Summarize the following text in \[NUMBER\] bullet points, focusing on
> \[FOCUS AREAS\] and ignoring \[EXCLUSIONS\]."

**Applied version:** \> "Summarize the following text in 3 bullet
points, focusing on main arguments and ignoring examples."

------------------------------------------------------------------------

## Test 1: Blog Post About Programming

### Input Text

Learning to code is like learning a new language. At first, the syntax
feels strange and the logic seems backward. But with consistent
practice, patterns emerge. You start recognizing common structures, and
what once took hours now takes minutes. The key is building projects,
not just completing tutorials. When you write real code to solve real
problems, the concepts stick. Many beginners give up too early,
expecting fluency overnight. But programming is a skill that compounds -
each hour invested makes the next hour more productive. Start small,
stay consistent, and trust the process.

### Prompt Used

> "Summarize the following text in 3 bullet points, focusing on main
> arguments and ignoring examples.
>
> \[Text inserted above\]"

### Model Output

-   Learning to code becomes easier with consistent practice as patterns
    and common structures become clearer over time.
-   Building real projects is more effective for learning than relying
    only on tutorials.
-   Programming skills compound with steady effort, so starting small
    and staying consistent leads to improvement.

------------------------------------------------------------------------

## Test 2: Technical Report

### Input Text

System Performance Analysis Report - Q4 2025

The database migration completed on December 15, 2025, resulting in a
40% reduction in query response times. Average latency decreased from
350ms to 210ms across all endpoints. CPU utilization dropped from 78% to
52% during peak hours. The new indexing strategy improved read
operations by 65% while write operations remained stable at baseline
performance. Three critical bugs were identified post-migration:
authentication token expiration (resolved Dec 18), cache invalidation
delays (resolved Dec 20), and connection pool exhaustion under load
(resolved Dec 22). Total system downtime during migration was 4.2 hours.
Rollback procedures were tested successfully but not needed.
Recommendation: Apply similar optimization to the analytics database in
Q1 2026.

### Prompt Used

> "Summarize the following text in 3 bullet points, focusing on main
> arguments and ignoring examples.
>
> \[Text inserted above\]"

### Model Output

-   The database migration significantly improved performance, reducing
    query response times and lowering latency and peak CPU utilization.
-   The new indexing strategy boosted read performance while maintaining
    stable write operations.
-   Post-migration issues were resolved and similar optimizations are
    recommended for the analytics database.

------------------------------------------------------------------------

## Analysis

### Pattern Consistency

The pattern worked consistently for both texts by providing exactly
three bullet points in both outputs. Despite one being a casual blog
post and the other a technical report, the template successfully
extracted the key information in the specified format.

### Improvement Over Weak Prompt

The pattern was able to point out in three bullet points the necessary
information, eliminating unnecessary details. The constraint on
"focusing on main arguments" forced the model to prioritize essential
information, while "ignoring examples" removed non-essential details
from the summaries.

### Reusability Factors

The pattern that makes it reusable includes the ability to adjust the
number of bullet points, what to focus on, what to ignore, and the text
itself. These four variables can be modified based on the use case:
technical documentation might need 5 bullet points focused on "metrics
and outcomes," while meeting notes might need 3 bullet points focused on
"action items" and ignoring "background discussion."

**Key Insight:** A well-designed prompt pattern transforms a vague
instruction into a flexible template that produces consistent,
high-quality outputs across diverse content types while remaining
adaptable to specific needs.