
The Apdex score (Application Performance Index) is calculated using the following formula:

***Apdex = (S + T/2) / N***

Where:
- S  is the number of **satisfied** requests (response times are below a threshold  T,
- T  is the number of **tolerating** requests (response times are between T and 4T,
- N is the total number of requests.

The Apdex score ranges from 0 to 1:
- 1 means all requests were satisfied.
- 0 means none of the requests were satisfied or tolerating.

Here’s how the score is broken down:
- Satisfied requests contribute fully (count as 1).
- Tolerating requests count as half (0.5).
- Dissatisfied requests (response time greater than 4T do not contribute to the score.

Example:
- 100 total requests ( N = 100 ),
- 70 are satisfied ( S = 70 ),
- 20 are tolerating ( T = 20 ),
- 10 are dissatisfied.


Apdex Score = (70 + 20/2) / 100 = 0.8

In this example, the Apdex score is 0.8.


### Why is dissatisified score not taken into account?


Dissatisfied requests are not factored directly into the Apdex score because the Apdex formula is designed to emphasize positive performance—meaning the focus is on measuring how well the application is performing for most users, not penalizing it for poor performance. The logic behind this decision includes:

Clarity in Feedback: The Apdex score is meant to provide a clear, concise picture of how many requests fall into the "satisfied" or "tolerating" categories, which gives a more optimistic and actionable view of performance.

Threshold-Based Evaluation: Dissatisfied requests have already exceeded the tolerance threshold (typically set at 4T, where 𝑇 is the acceptable response time). Since these requests are considered failures, they don't contribute positively to the score.

Simplified Weighting: By excluding dissatisfied requests, Apdex focuses on ranking the overall user experience between fully satisfied and tolerable. Dissatisfied requests don’t contribute positively because they’re treated as failures or negative experiences.

However, dissatisfied requests do affect the score indirectly:

They lower the number of satisfied or tolerating requests, thus reducing the score.
