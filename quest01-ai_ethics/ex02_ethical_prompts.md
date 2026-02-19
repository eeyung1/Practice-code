# Exercise 2: Designing Ethical Prompts

---

## Part A: Independent Analysis

### Prompt #1: Gender/Profession Bias

**Bad Prompt:**
> "What's the country with the best doctors"

**Why it's problematic:**
There's no specificity for what age or criteria defines "best." It reduces rich medical traditions and individual excellence to national competition, creates artificial hierarchies between countries, and assumes that "best" is universal across all medical specialties when medicine is highly diverse.

**Potential harm:**
Most people would avoid their country's doctors to seek medical care elsewhere, like many Nigerians seeking medical procedures in the Western world. This creates brain drain, dismisses local expertise, and can lead to unnecessary medical tourism based on bias rather than actual quality of care.

**Hidden assumptions:**
The prompt assumes "best" means the same thing across cultures when healthcare priorities differ globally (preventative vs curative care). It ignores that medical training is increasingly global, and assumes that best doctors come from countries with the best economies rather than recognizing individual skill and training quality.

**Improved Prompt:**
> "What are some leading cardiology research institutions globally and what breakthroughs has each contributed?"

---

### Prompt #2: Cultural/Regional Bias

**Bad Prompt:**
> "Which country in Africa has the best jollof rice"

**Why it's problematic:**
The definition of "best" is not defined when it comes to food which is subjective. It reduces rich cultural culinary traditions to competition and creates artificial hierarchies between neighboring West African cultures.

**Potential harm:**
It creates tensions between countries like the ongoing Nigerians versus Ghanaians online debate about jollof rice. This reinforces nationalist rivalries, can affect cultural exchange and tourism, and dismisses the culinary traditions of other African countries entirely.

**Hidden assumptions:**
Best is subjective and varies within countries and across different taste preferences. The prompt assumes jollof rice is the same dish across regions when it actually uses different spices, cooking methods, and ingredients depending on location. It ignores that "best" depends entirely on who's judging — diaspora versus locals, different regional preferences within countries.

**Improved Prompt:**
> "What are the unique regional variations of jollof rice across West Africa and what makes each version special?"

---

### Prompt #3: Age/Generation Bias

**Bad Prompt:**
> "Are older people bad with technology"

**Why it's problematic:**
There's no specificity for what age counts as "older" — is it 40, 60, or 80? It treats all older people as a monolithic group while ignoring massive individual differences in skill and experience. The prompt creates a false binary between good and bad with technology when ability exists on a spectrum.

**Potential harm:**
In a world greatly influenced by technology there would be hiring discrimination against older workers who are perfectly qualified. This forces experienced workers out of jobs, creates a self-fulfilling prophecy where older people avoid learning technology because they believe the stereotype about themselves, and wastes valuable institutional knowledge and experience.

**Hidden assumptions:**
The prompt assumes technology is one monolithic thing when using TikTok versus programming databases are completely different skills. It ignores that people who are 60+ today literally invented personal computers, the internet, and smartphones. It also assumes young automatically equals tech-savvy when many young people only know how to use apps but don't understand underlying systems.

**Improved Prompt:**
> "What are examples of successful content creators across different age groups and what skills do they share?"

---

## Part B: Test and Validate

### Testing Prompt #1

**Original Bad Prompt Response:**
> "There isn't a single country with the 'best' doctors — medical quality depends on training standards, healthcare systems, and specialties. Countries frequently recognized for excellent medical training and healthcare systems include: United States, Germany, Switzerland, Japan, South Korea. 'Best' often depends on the medical field (e.g., surgery, oncology, cardiology) and access to advanced technology and research."

**Improved Prompt Response:**
> "Leading cardiology research institutions globally have made major contributions to modern cardiovascular medicine. In the United States, institutions such as Cleveland Clinic, Mayo Clinic, and Johns Hopkins Hospital have driven advances in coronary artery bypass surgery, artificial heart devices, cardiac imaging, electrophysiology, and research linking cholesterol and inflammation to heart disease. These centers have helped shape global treatment guidelines and improved survival rates for heart patients.
>
> In Europe, Royal Brompton Hospital in the United Kingdom has led research in congenital heart disease and pulmonary hypertension, while Deutsches Herzzentrum München in Germany has pioneered heart transplantation programs and mechanical circulatory support technologies. These institutions have contributed significantly to both surgical innovation and pediatric cardiology.
>
> In Asia, National Cerebral and Cardiovascular Center has advanced hypertension research and stroke prevention, and Fuwai Hospital has conducted large-scale cardiovascular epidemiology studies and improved interventional cardiology techniques. Collectively, these institutions have played critical roles in breakthroughs ranging from bypass surgery and heart transplants to advanced imaging, stent technology, and preventive cardiology."

**Analysis:**
Even though AI tried to be diplomatic with the bad prompt, it was still forced into problematic territory by listing countries and creating an implicit ranking. The improved prompt guided AI to focus on specific institutions and their unique contributions rather than national hierarchies, resulting in an educational and non-biased response.

---

### Testing Prompt #2

**Original Bad Prompt Response:**
> "There is no objectively 'best' jollof rice. Taste is subjective, and different countries have their own preferred styles and flavors."

**Improved Prompt Response:**
> "West African jollof rice varies by region. Nigeria often features smoky flavor and bold seasoning, Ghana is known for fragrant basmati rice and rich tomato base, and Senegal (home of thieboudienne) traditionally includes fish and vegetables cooked together. Each version reflects local ingredients and cultural preferences."

**Analysis:**
The bad prompt forced AI into a defensive position where it had to debunk the premise rather than provide useful information. The improved prompt led to a celebratory response that honored each region's unique culinary traditions and provided educational value about cultural diversity.

---

### Testing Prompt #3

**Original Bad Prompt Response:**
> "No. Technology ability depends on exposure, training, and interest—not age. Many older adults use digital tools effectively."

**Improved Prompt Response:**
> "Examples include younger creators like MrBeast and older creators like Grandma Moses (who began her career later in life). Across age groups, successful creators share skills such as creativity, consistency, adaptability, storytelling, and audience engagement."

**Analysis:**
The bad prompt required AI to waste energy debunking a stereotype before providing any useful information. The improved prompt immediately delivered concrete examples and actionable insights about skills that matter regardless of age, making the response both inspiring and educational.

---

### AI Critique of Improved Prompts

**Question asked to AI:**
> "I rewrote three biased prompts to be more ethical: [listed all three]. My reasoning: I wanted to avoid rankings, celebrate diversity, and focus on specific contributions rather than broad generalizations. What ethical issues might I have still missed? How could these prompts still go wrong?"

**AI Response:**
Your revisions are strong, but a few ethical risks remain:

**1. Hidden Framing Bias**
Even without rankings, prompts may still imply hierarchy (e.g., "leading" institutions privileges Global North research centers and underrepresents lower-income regions).

**2. Selection Bias**
Asking for "successful" examples may highlight highly visible figures while ignoring structural barriers that limit visibility for others.

**3. Representation Gaps**
"Global" or "across West Africa" can unintentionally exclude smaller countries or minority variations if the answer focuses only on the most documented cases.

**4. Positive Framing Bias**
Avoiding stereotypes is good, but focusing only on success stories can oversimplify reality and hide systemic inequalities.

**5. Cultural Essentialism Risk**
Even improved cultural prompts (e.g., jollof rice variations) can still unintentionally treat regions as homogeneous or fixed in identity.

**Overall:** You reduced explicit bias (rankings and stereotypes), but risks remain around representation, power imbalance, framing effects, and whose stories are considered "notable."

---

## Part C: Reflection

### If I'd asked AI to write ethical prompts from the start, would I understand why they're ethical?

If I asked AI to write ethical prompts I wouldn't have understood why they're ethical, considering the original prompts would have been biased in the first place. But after struggling to write bad prompts and then analyzing them deeply — identifying why they're problematic, what harm they cause, and what assumptions they contain — I gained firsthand knowledge of what makes a prompt ethical or unethical. Testing both versions with real AI and seeing the difference in outputs made the lesson concrete rather than theoretical.

### Can I now design ethical prompts without AI?

Yes. Considering what I have learned so far, I can design ethical prompts that avoid rankings and hierarchies, celebrate diversity rather than competition, and focus on specifics rather than broad generalizations. I now understand that even well-intentioned prompts can carry hidden biases through framing, selection, and representation gaps. The key is to ask myself: does this prompt force comparison, does it assume universality, and whose voices might it exclude?
