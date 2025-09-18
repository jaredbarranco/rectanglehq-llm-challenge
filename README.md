# Rectangle HQ Application Repository
## Intro
I had a great conversation with Marvin Kennis at RectangleHQ about integration challenges in the logistics industry. Since then, I've kept an eye out on Rectangle's job openings. They recently posted a position for their first Forward Deployed Engineer.

Their application starts with a POST request with a header... very cool and drew my attention immediately. I am not deluding myself into thinking that I am a top candidate for this role. However, given my breif personal conversation with Marvin and my interest in the way they presented this problem, I wanted to give it a go.

I had two thoughts on how to stand out.

1. Build a script to do this for me, effectively in a "Notebook" style progression.

2. Use an LLM to apply, since each step of the application likely would continue requiring additional data, and potentially more interesting header manipulation. Plus, I wouldn't mind this as a challenge.

### LLM Based Solution

I started with the LLM approach, as I wanted to experiment with sandboxing an LLM using containers. We are looking at implementing team wide LLM usage at my current employer, so this seemed like a relevant task. Additionally, since this would be making network requests, I didn't want it to have full file access to my host machine.

I used ChatGPT to create a Containerfile which configured: non-root user on an Ubuntu image, installed some of my preferred tools.

I then just described the challenge to Google's Gemini 2.5 Pro and got underway.

Some interesting challenges: 
1. I had was that the agent had a difficult time with decoding the data because the curl progress report interrupted the stream on stdout. In the future, if I'm running `curl` with agents, I'll hide the progress report.

2. I'd never dealt with decoding hex encoded values, so I just wrote the encoded string to a file and had the LLM determine what the best course of action was.

3. The container image didn't have some utilities setup - primarily no EDITOR or buffer handling. This meant anytime Opencode tried to write something to my clipboard, nothing was copied to the system buffers. Additionally, if I tried to open things in a default editor, nothing would happen. I ended up correcting these outside of my conversation thread.

4. The final payload required me to include my name and email in the payload again, even though the submission instructions explicitly said: `Submit the optimal solution as {code, answer}`. This threw me for a loop when submitting, as I performed the final submission `curl` call myself.

My takeaways here are:

1. Overall, the LLM excelled at this task - because it effectively ended up being a Dynamic Programming Leetcode problem behind some curl calls. I really enjoyed the process of configuring the LLM to attack this problem, and working alongside it to unblock it from the limitations I placed on it. 

2. LLM Sandboxing seems like it could be beneficial and do-able for a small team. Obviously the concept of dev containers aren't new, but this is more pointed at security concerns, and less about environmental setup.

3. I can acknowledge that my submission probably won't be taken seriously. As you can see in the optimize.go file in the repo, I called out my use of LLM assistance and directed them to this README.


### Notebook Style Progression
I may re-approach this another time and implement a true Jupyter Notebook type implementation. I am confident in my ability to use and apply new technologies and skills to difficult problem domains, and this would be no different.

