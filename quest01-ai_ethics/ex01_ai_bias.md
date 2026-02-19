# Exercise 1: Understanding AI Bias

---

## Part A: Build Your Foundation (No AI First)

### Bias Example

When I asked AI to give the name of the country that comes to mind when talking about online fraud, Nigeria was generated even though statistically it is not ranked within the top 3. I consider that a bias against the citizens of the country.

### Why This Bias Exists

Most of the data used in AI training has been highly influenced by the Western world since there is more content from them on the internet than from countries like Nigeria. When it comes to stories about Nigeria online, the stories that get amplified are the fraud stories, which has even resulted in PayPal suspending Nigerians from receiving payments. This affected so many freelancers and businesses in Nigeria. That's a major bias with real-world consequences.

### Proposed Mitigation Strategy

I propose two mitigation strategies:

**1. Balance Training Data**
Since AI relies heavily on data, having balanced training data would ensure any response generated would have a certain degree of accuracy and reduces bias. This means actively including more positive Nigerian content — tech achievements, cultural contributions, scientific innovations — rather than overrepresenting sensationalized fraud stories.

**2. Fact-Checking Layer**
During the fact-checking layer, the AI should compare the statistics available on the countries associated with online fraud so as to be able to bring out the number one on the list instead of generating Nigeria due to social bias. This would prevent AI from relying purely on frequency of mention in text and instead ground responses in verifiable data.

---

## Part B: Strategic AI Use

### Question Asked to AI

> "I found this bias example: AI associates Nigeria with online fraud despite Nigeria not being statistically in the top 3 countries for online fraud.
>
> I think it exists because Western-dominated training data amplifies negative fraud stories over positive Nigerian tech achievements, leading to real consequences like PayPal bans.
>
> My proposed mitigation: (1) Balance training data with more diverse Nigerian content, and (2) Add fact-checking layer that verifies statistics before generating responses.
>
> What am I missing? What underlying causes haven't I considered?"

### AI Response

The bias associating Nigeria with online fraud does not arise only from unbalanced training data, but from deeper structural and historical factors. Long-standing stereotypes, such as the globally popularized "Nigerian prince scam," have shaped internet culture and media narratives for decades. AI systems trained on large-scale internet text inevitably absorb these narratives. As a result, models may reproduce historical cultural bias even if statistical evidence does not support the association.

Another underlying cause is frequency and salience bias. AI models learn from patterns of co-occurrence in text, not from verified crime statistics. If sensational fraud stories mentioning Nigeria appear more frequently or receive more global attention than stories about Nigerian innovation, technology startups, or academic achievements, the model will disproportionately associate Nigeria with fraud. This reflects how media amplification and virality shape machine learning outputs.

There are also systemic and geopolitical factors at play. Much of the world's digital content, major media platforms, and AI development institutions are concentrated in Western countries. This creates epistemic imbalance, where certain regions dominate global narratives while others are underrepresented. Additionally, platform-level systems such as fraud detection algorithms and risk-scoring tools can reinforce bias through feedback loops: increased scrutiny leads to more reports, which generate more data reinforcing the stereotype.

Technically, proxy variable bias can also contribute to the issue. Even if the word "Nigeria" is not directly used negatively, correlated signals such as geolocation data, linguistic patterns, or payment routing information may function as indirect indicators. Furthermore, language models do not automatically normalize statistics or compare base rates, meaning they may generate responses based on common associations rather than contextualized facts.

To address these deeper causes, mitigation should go beyond simply adding more Nigerian content. It should include counterfactual data balancing that highlights innovation and positive contributions, fairness auditing before deployment, contextualized response framing that presents proportional statistics, and human oversight for sensitive outputs. Ultimately, this issue reflects broader social, economic, and narrative power structures that AI systems inherit and amplify rather than create independently.

### Real-World Implications in Three Scenarios

**1. Hiring Decisions (AI Resume Screening)**

When bias associating Nigeria with fraud appears in AI-driven hiring systems, it can result in systemic employment discrimination. Automated resume screening tools may rank candidates lower if they detect Nigerian educational institutions, work experience, or location indicators. Even without explicit intent, the system may treat these signals as higher risk due to learned associations. Because hiring algorithms operate at scale, many qualified applicants could be filtered out before a human ever reviews their applications.

The real-world implication is reduced access to employment opportunities for innocent individuals. This limits career mobility, suppresses income potential, and may push talented professionals to conceal aspects of their background. Over time, such bias can contribute to brain drain and reduce trust in digital recruitment systems.

**2. Financial Services (Loans, Payments, Fraud Detection)**

In financial services, biased AI systems can cause financial exclusion. Risk-scoring algorithms may flag Nigerian applicants as higher fraud risks, leading to loan denials, frozen accounts, payment restrictions, or higher interest rates. These outcomes may occur even when individuals have strong credit histories or legitimate businesses.

The broader implication is restricted access to capital and reduced participation in the global economy. Entrepreneurs may struggle to grow businesses, individuals may lack access to credit for education or housing, and digital trade opportunities may shrink. This creates a harmful feedback loop where financial exclusion reinforces economic inequality.

**3. Immigration and Visa Processing**

If bias enters immigration or visa decision systems, Nigerian applicants may face higher rejection rates, longer processing times, or increased documentation requirements. Automated risk assessment tools could disproportionately classify applicants as high risk based on nationality-linked patterns rather than individual circumstances.

The real-world consequences include limited global mobility, reduced access to international education and employment, family separation, and institutionalized discrimination. When governments rely on algorithmic systems, bias can become embedded within official decision-making processes, making it harder to detect and challenge.

**Overall Impact:**

The most serious implication across all three areas is scalable discrimination presented as objective and data-driven. Algorithmic bias can operate quickly, affect large populations simultaneously, and reinforce existing social and economic inequalities. Without fairness auditing, transparency, and accountability, such systems risk amplifying historical stereotypes rather than promoting equitable outcomes.

---

## Part C: Critical Reflection

### What percentage did I complete before using AI?

I completed approximately 70% before using AI. I identified the bias example myself, explained why it exists based on my understanding of Western-dominated training data and media amplification, and proposed two concrete mitigation strategies. AI then deepened my understanding by introducing new concepts and systemic factors I hadn't considered.

### Did AI replace my thinking or amplify it?

AI amplified my thinking by providing detailed analysis of my initial reasoning and introducing new concepts like salience bias, feedback loops, epistemic imbalance, and proxy variable bias. These concepts built on top of what I already understood rather than replacing my thinking. I learned how the bias operates at multiple levels — not just in training data, but through historical narratives, media systems, and institutional processes.

### Could I explain this to someone else without AI?

Yes. I could explain that AI bias against Nigeria exists because training data reflects Western-dominated internet content where sensational fraud stories get amplified over positive tech achievements. This creates a feedback loop where stereotypes become self-reinforcing through media attention and platform policies. The bias has real consequences like PayPal bans, hiring discrimination, and financial exclusion. I would explain that mitigation requires both balanced training data and fact-checking layers to ground AI responses in statistics rather than cultural stereotypes.

### What did I contribute that AI couldn't?

The Nigeria fraud example came from my lived experience as a Nigerian who has seen this bias firsthand. The PayPal harm example came from my knowledge of how this bias has affected Nigerian freelancers and businesses in real, tangible ways. AI can explain bias theoretically, but I brought the personal relevance and real-world context of someone directly impacted by this specific bias. This perspective — understanding bias not as an abstract concept but as something that blocks economic opportunities for real people — is something AI cannot generate on its own.
