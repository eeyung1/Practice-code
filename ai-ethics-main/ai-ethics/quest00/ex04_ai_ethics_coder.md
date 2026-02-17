## Exercise 4 — Ai as a Learning Amplifier Goal

**Use Ai the right way**
 think first, struggle first, then use AI to deepen understanding — not replace it.

# Topic: Wireless Routing Protocols

Phase 1  Building My Own Foundation (Before Ai)
 1. # Active Reading

**I started by reading about basic wireless routing protocols such as**

Distance Vector Routing

Link State Routing

**AODV (Ad hoc On-Demand Distance Vector)**

**OLSR (Optimized Link State Routing)**

After reading, I closed the material and tried to explain it in my own words.

2. # Self-Explanation (Without Looking)

Wireless routing protocols help devices communicate in networks where there is no fixed infrastructure.

Some protocols constantly share routing tables (proactive), while others only find routes when needed (reactive).

Proactive protocols (like OLSR) maintain updated routes all the time.
→ Faster communication, but more overhead.

Reactive protocols (like AODV) find routes only when necessary.
→ Less overhead, but possible delay when sending data.

The key trade-off is
Speed vs. Network Overhead.

3. # Small Scenario (5–10 Devices)

I designed a small network of

8 wireless sensors

1 central monitoring device

**Scenario** Environmental monitoring in a small warehouse.

Devices are spread across the space and communicate wirelessly.

**My reasoning**

Since the network is small and stable, proactive routing might work well.

Devices don’t move often.

Low latency is important for alerts.

So I would choose a proactive protocol like OLSR.

Justification:

Small network → overhead is manageable.

Stable topology → routing tables won’t change often.

Faster response time for alerts.

# Phase 2 — Strategic AI Use

After building my own understanding, I would ask AI targeted questions like

“In what situations would reactive routing outperform proactive routing?”

“How does mobility affect AODV performance?”

“What are the scalability limits of OLSR?”

“How do wireless routing protocols handle packet loss?”

This helps me test my mental model.

For example, I learned:

Reactive protocols are better for highly dynamic networks.

Proactive protocols struggle when the network grows large.

Energy consumption becomes a critical factor in sensor networks.

Ai did not replace my thinking. It stress-tested it.

# Phase 3 — Real Application
**Smart City Network Design**
Scenario

1,000 IoT sensors (air quality, temperature, parking sensors)

50 traffic lights

10 emergency vehicles

# Step 1 My Initial Design (Before AI Feedback)
IoT Sensors (1,000 devices)

These are low-power and mostly static.

I would likely use

A reactive or hybrid routing protocol.

Possibly hierarchical clustering to reduce overhead.

**Reason**

1,000 devices is large.

Proactive routing would create too much constant traffic.

Energy efficiency is critical.

Traffic Lights (50 devices)

These need high reliability and low latency.

I would use

A more stable, possibly proactive approach.

Possibly centralized coordination.

Reason

Traffic lights must synchronize.

Delays could cause accidents.

Emergency Vehicles (10 devices, mobile)

Highly mobile nodes.

**I would use**

A reactive protocol like AODV.

**Reason**

Routes must adapt quickly as vehicles move.

Flexibility is more important than constant routing updates.

Failure Points I Identified

Network congestion due to 1,000 sensors.

Packet collision in dense areas.

Battery drain in IoT nodes.

Route instability with moving emergency vehicles.

Security vulnerabilities (spoofing, interference).

After AI Feedback

# Ai might suggest

Using hybrid protocols like ZRP (Zone Routing Protocol).

Using mesh networking for redundancy.

Implementing QoS prioritization for emergency vehicles.

Segmenting the network into layers (sensor layer, control layer, emergency layer).

This refinement helped me think more systematically about scale and reliability.

# Reflection
% Human vs Ai Contribution

**I would say**

70% human reasoning

30% Ai refinement

I built the structure and core decisions myself.
Ai helped things i colude expand edge cases and optimization strategies.

**Could I Defend My Decisions Without Ai?**

Yes.
I understand

Proactive vs reactive trade-offs.

Scalability challenges

Energy constraints

Mobility effects

Even without Ai, I could explain my logic clearly.

**What Will I Remember in 6 Months?**

The proactive vs reactive trade-off.

That scalability changes everything.

That mobility affects routing stability.

That energy efficiency is critical in IoT systems.

That design always involves trade-offs.

**Did Ai Make Me Sharper or Think for Me?**

# Ai made me sharper because

I struggled first

I built my own mental model

I used AI to test and improve it

If I had asked Ai to “design a smart city network” from the start, I would have copied something impressive but shallow in understanding.

**This exercise showed me something important**

Ai is powerful — but it only strengthens me if I think first.

Otherwise, it weakens my independence that is why i must understand the way i use ai.
