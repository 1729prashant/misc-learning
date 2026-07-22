## PeopleCode traversal
```
/* Declaration */

Local Rowset &rs_level0, &rs_level1, &rs_level2, &rs_level3;
Local Row &row_level0, &row_level1, &row_level2, &row_level3;
Local Record &rec1;

/* PROCESSING LEVEL 0 */

&rs_level0 = GetLevel0();
&row_level0 = &rs_level0.GetRow(1);

/* PROCESSING LEVEL 1 */

&rs_level1 = &row_level0.GetRowset(Scroll.LEV1COUNTRY_TBL);
For &i = 1 To &rs_level1.ActiveRowCount
   &row_level1 = &rs_level1(&i);
   
   
   /* PROCESSING LEVEL 2 */

   &rs_level2 = &row_level1.GetRowset(Scroll.LEV2_STATE_TBL);
   For &j = 1 To &rs_level2.ActiveRowCount
      &row_level2 = &rs_level2(&j);
      
      
      /* PROCESSING LEVEL 3 */

      &rs_level3 = &row_level2.GetRowset(Scroll.LEV3_DIST_TBL);
      For &k = 1 To &rs_level3.ActiveRowCount
         &row_level3 = &rs_level3(&k);
         
         /*GET RECORD*/
         
         &rec1 = &row_level3.LEV3_DIST_TBL;
         
         /*GET FIELD*/
         
         &FIELD1 = &rec1.COUNTRY_DISTNAME.Value;
         WinMessage("District Name Is" | &FIELD1);
         
      End-For;
   End-For;
End-For;

```


## Unix search
```
--recursively find all files in current and subfolders, find needs a starting point, so 
--the . (dot) points to the current directory
find . -name "foo*"

--case insensitive
find . -iname "foo*" 

--recursive search for text in files in a folder using grep
grep -Rnw '/path/to/somewhere/' -e 'pattern'
-- -r or -R is recursive ; use -R to search entirely
-- -n is line number, and
-- -w stands for match the whole word.
-- -l (lower-case L) can be added to just give the file name of matching files.
-- -e is the pattern used during the search
-- r option is lazy (traverses depth-first, than stops after the first directory)
-- R is greedy (will traverse the entire tree correctly).

--Along with these, --exclude, --include, --exclude-dir flags could be used for efficient 
--searching:
--This will only search through those files which have .c or .h extensions:
grep --include=\*.{c,h} -rnw '/path/to/somewhere/' -e "pattern"

-- This will exclude searching all the files ending with .o extension:
grep --exclude=\*.o -rnw '/path/to/somewhere/' -e "pattern"

--For directories it's possible to exclude one or more directories using the --exclude-dir 
--parameter. For example, this will exclude the dirs dir1/, dir2/ and all of them matching 
--*.dst/:
grep --exclude-dir={dir1,dir2,*.dst} -rnw '/path/to/search/' -e "pattern"

-- source:
https://stackoverflow.com/questions/16956810/find-all-files-containing-a-specific-text-string-on-linux

```

## Misc Unix

Description  | Link
------------- | -------------
Vim cheat sheet  | <https://vim.rtorr.com>
GNU grep cheat sheet  | <https://devhints.io/grep>
Bash scripting cheat sheet | <https://devhints.io/bash>


## Flowdiagram, drawing Tools

Description | Link
------------- | -------------
Excalidraw | <https://excalidraw.com>
Figma | <https://www.figma.com>
Mural | <https://app.mural.co>
Eraser | <https://app.eraser.io/dashboard/all>
Draw.io | <https://www.drawio.com>



## UX/UI

Description | Link
------------- | -------------
Storyboard Website Design — How to Pre-Visualize Your Site | <https://www.studiobinder.com/blog/storyboard-website-design/>
What Is UI Design? Definition, Tips, Best Practices | <https://www.coursera.org/articles/ui-design>
UI Design, Adobe | <https://xd.adobe.com/ideas/process/ui-design/>
User Interface (UI) Design, Interactive Foundation Design| <https://www.interaction-design.org/literature/topics/ui-design>
Usability, Digital.gov USA| [Usablility](https://digital.gov/topics/usability/#:~:text=User%20research%20focuses%20on%20understanding,of%20design%20on%20an%20audience.%E2%80%9D)
Scenario Mapping: Design Ideation Using Personas | <https://www.nngroup.com/articles/scenario-mapping-personas/>
User scenarios | <https://guides.18f.gov/methods/decide/user-scenarios/>
10 most popular design systems to learn from in 2022 for UX Designers | <https://uxplanet.org/10-most-popular-design-systems-to-learn-from-in-2022-for-ux-designers-18a24843a860>
Design Systems 101 | <https://www.nngroup.com/articles/design-systems-101/>
16 Tips that Will Improve Any Online Form | <https://uxplanet.org/the-18-must-do-principles-in-the-form-design-fe89d0127c92>
31 Examples Of Icons In Navigation Menus | <https://www.awwwards.com/31-examples-of-icons-in-navegation-menus.html>
7 Principles of Icon Design | <https://uxdesign.cc/7-principles-of-icon-design-e7187539e4a2?gi=7fac793515a7>
Usability Evaluation | <https://www.interaction-design.org/literature/book/the-encyclopedia-of-human-computer-interaction-2nd-ed/usability-evaluation>
Interaction Design Evaluation Methods | <https://gayan1999malinda.medium.com/interaction-design-evaluation-methods-df8132cedbf9>
UI Code Examples - uiverse.io| <https://uiverse.io/>


## Deception, Counter Deception, Critical Thinking

|No.    |MAXIM                                    |COUNTER MAXIM                            |
|-------|-----------------------------------------|-----------------------------------------|
|1.     |Magruder’s Principle: It is easier to maintain a preexisting belief in the target than force a change in beliefs. |Apply the same critical analysis to facts that support your assumptions that you apply to facts that contradict them.            |
|2.     |Exploiting Limits of Human and Machine Sensing and Information Processing. |Develop lots of trustworthy sources of data and analysis. Measure reality from multiple perspectives and resolutions.            |
|3.     |Jones Dilemma: Deception is difficult if there are more real sources than false sources. |Develop lots of trustworthy sources of data and analysis. Measure reality from multiple perspectives and resolutions.|
|4.     |Carefully Sequence Deception Activities to Tell a Story. |Is there an "orgy of evidence?" Are facts being revealed all at once when it would make more sense for them to emerge slowly (or vice-versa).|
|5.     |Carefully Design Planned Placement of Deceptive Material. Make target “work” for it, don’t boldly announce what you are doing (example diary easily found on a desk vs. hidden, obfuscated code). Contrast this with “orgy of evidence”. |Don't assume a fact is true just because it was hard to obtain it.|
|6a.    |Ambiguity - Increase doubt by providing many possible truths. |Is there more than one narrative present? Contradictions and incongruities may suggest deception. |
|6b.    |Misdirection Decrease doubt by focusing target on a given falsehood. |The Plus-Minus Rule: No imitation can be perfect without being the real thing. In a simulation, characteristics will be added or they will be missing. |
|7.     |Husband Deception Assets. |Model the adversary's capabilities and neutralize them. |
|8.     |Feedback - Includes attackers monitoring for success of deception and monitoring for being deceived themselves. |Behave as if you believe the deception so the adversary doesn't pivot. Deter the adversary by showing them that you've detected the deception. (Pre-bunking) |




[DEF CON 32 - Tom Cross Greg Conti - Deception & Counter Deception - Defending Yourself in a World Full of Lies.pdf
](https://media.defcon.org/DEF%20CON%2032/DEF%20CON%2032%20presentations/DEF%20CON%2032%20-%20Tom%20Cross%20Greg%20Conti%20-%20Deception%20%26%20Counter%20Deception%20-%20Defending%20Yourself%20in%20a%20World%20Full%20of%20Lies.pdf)

Below links are from above DEFCON talk, added for redundancy only

##### Deception:
* [CIA Deception Maxims](https://www.governmentattic.org/18docs/CIAdeceptionMaximsFactFolklore_1980.pdf)
* [Army Support to Military Deception](https://armypubs.army.mil/epubs/DR_pubs/DR_a/ARN15310-FM_3-13.4-000-WEB-2.pdf)
* [US DoD Doctrine on Military Deception](https://info.publicintelligence.net/JCS-MILDEC.pdf)

##### Deception in Malware:
* [The Untold Story of the 2018 Olympics Cyberattack, the Most Deceptive Hack in History](https://www.wired.com/story/untold-story-2018-olympics-destroyer-cyberattack/)
* [OlympicDestroyer is here to trick the industry](https://securelist.com/olympicdestroyer-is-here-to-trick-the-industry/84295/)
* [Under False Flag: Using Technical Artifacts for Cyber Attack Attribution](https://cybersecurity.springeropen.com/articles/10.1186/s42400-020-00048-4)
* [Wave Your False Flags! Deception Tactics Muddying Attribution in Targeted Attacks](https://media.kasperskycontenthub.com/wp-content/uploads/sites/43/2017/10/20114955/Bartholomew-GuerreroSaade-VB2016.pdf)
* [Digital False Flag Operations: A How-To Guide](https://grimminck.medium.com/digital-false-flag-operations-a-how-to-guide-bc529b54cc22)

##### Counter-Deception:
* [Counterdeception Principles and Applications for National](https://www.amazon.com/Counterdeception-Principles-Applications-National-Security/dp/1580539351)

##### Security Journalism, Media & Critical Thinking:
* [48 Questions for Critical Thinking - Justin Wright](https://www.linkedin.com/posts/jwmba_90-of-people-lack-critical-thinking-skills-activity-7193961300748562432-exPW/?utm_source=share&utm_medium=member_desktop)
* [Media Literacy Now](https://medialiteracynow.org/)
* [Journalism as a scientific endeavor - Julia Angwin](https://www.proofnews.org/a-letter-from-our-founder/)
* [How do we protect ourselves against disinformation?- Tom Cross](https://medium.com/@_decius_/how-do-we-protect-ourselves-against-disinformation-88f798b0d260)
* [Facts, frames, and (mis)interpretations: Understanding rumors as collective sensemaking – Kate Starbird](https://www.cip.uw.edu/2023/12/06/rumors-collective-sensemaking-kate-starbird/)

##### Internet Counter-Deception:
* [DISARM Framework](https://www.disarm.foundation/framework)
* [Misinformation Village](https://www.misinfovillage.org/)
* [RAND: Tools that Fight Disinformation Online](https://www.rand.org/research/projects/truth-decay/fighting-disinformation/search.html)
* [The Socratic Web](https://medium.com/@Aegist/what-is-the-socratic-web-c6095c452c6) & [Why Tools Shouldn't Adjudicate Truth - Shane Greenup](https://medium.com/@Aegist/we-wont-solve-the-misinformation-problem-with-systems-which-adjudicate-on-truth-or-on-who-to-trust-bb02c835395)

##### Wikipedia:
* [WikiScanner](https://en.wikipedia.org/wiki/WikiScanner)
* [WikiWatchdog](https://github.com/volpino/WikiWatchdog)
* [Puppy Smoothies (by Tom Cross)](https://firstmonday.org/ojs/index.php/fm/article/view/1400/1318)
* [WikiTrust](https://en.wikipedia.org/wiki/WikiTrust)

##### Inspiration for the World Wide Web:
* [As We May Think - Vannevar Bush](https://www.theatlantic.com/magazine/archive/1945/07/as-we-may-think/303881/)
* [Augmenting Human Intellect – A Conceptual Framework - Douglas Engelbart](https://www.dougengelbart.org/pubs/augment-3906.html)
* [Hyperscope 2.0](https://dougengelbart.org/content/view/355/)
* [Computer Lib/Dream Machines - Ted Nelson](https://computerlibbook.com/)
* [Project Xanadu](https://www.xanadu.net/)


## Critical Thinking
[*(From 48 Questions for Critical Thinking)*](https://www.linkedin.com/posts/infographic-insights_48-questions-for-critical-thinking-credit-activity-7199378592881934338-FCD2/)

|||
|----------|----------|
| **Who** *<br>1. Who is affected by this issue ?<br>2. Who faces the biggest consequences ?<br>3. Who holds the power in this situation ?<br>4. Who might see this differently ?<br>5. Who are the key stakeholders involved ?<br>6. Who benefits from this outcome ?<br>7. Who else should be consulted ?<br>8. Who can provide more information ?* | **What** *<br>1. What is the issue at hand ?<br>2. What are the main arguments ?<br>3. What is the evidence ?<br>4. What assumptions are being made ?<br>5. What are the potential consequences ?<br>6. What alternatives exist ?<br>7. What are the risks of each alternative ?<br>8. What steps can be taken next ?* |
| **Where** *<br>1. Where did this first become an issue ?<br>2. Where is the problem most evident ?<br>3. Where can we find supporting data ?<br>4. Where have solutions worked before ?<br>5. Where are resources most needed ?<br>6. Where are potential obstacles located ?<br>7. Where can we implement solutions first ?<br>8. Where should we monitor the outcomes ?* | **When** *<br>1. When did the issue first emerge ?<br>2. When do the effects typically appear ?<br>3. When was the data last collected ?<br>4. When is the best time to act ?<br>5. When have solutions been attempted ?<br>6. When is the deadline for action ?<br>7. When should we expect to see results ?<br>8. When will we review progress ?* |
| **Why** *<br>1. Why is this issue significant ?<br>2. Why did it arise in the first place ?<br>3. Why are certain solutions preferred ?<br>4. Why might opinions differ ?<br>5. Why has this not been addressed sooner ?<br>6. Why are some more affected than others ?<br>7. Why is immediate action necessary ?<br>8. Why should we revisit in the future ?* | **How** *<br>1. How did this issue start ?<br>2. How does this impact different groups ?<br>3. How can we gather more data ?<br>4. How have others tackled similar issues ?<br>5. How will we implement the solution ?<br>6. How will we measure success ?<br>7. How should we communicate changes ?<br>8. How often should we reassess ?* |


From [*(Questions for a Socratic Dialogue)*](https://courses.cs.vt.edu/cs2104/Summer2014/Notes/SocraticQ.pdf) and [www.criticalthinking.org](https://www.criticalthinking.org/)

| Category | Questions |
|----------|-----------|
| **I. Questions of Clarification** | 1. What do you mean by ___ ?<br>2. What is your main point ___ ?<br>3. How does ___ relate to ___ ?<br>4. Could you put that another way?<br>5. What do you think is the main issue here?<br>6. Is your basic point ___ or ___ ?<br>7. Could you give me an example?<br>8. Would this be an example: ___ ?<br>9. Could you explain that further?<br>10. Would you say more about that?<br>11. Why do you say that?<br>12. Let me see if I understand you; do you mean ___ or ___ ?<br>13. How does this relate to our discussion/problem/issue?<br>14. What do you think John meant by his remark? What did you take John to mean?<br>15. Jane, would you summarize in your own words what Richard has said? Richard, is that what you meant? |
| **II. Questions That Probe Purpose** | 1. What is the purpose of ___ ?<br>2. What was your purpose when you said ___ ?<br>3. How do the purposes of these two people vary?<br>4. How do the purposes of these two groups vary?<br>5. What is the purpose of the main character in this story?<br>6. How did the purpose of this character change during the story?<br>7. Was this purpose justifiable?<br>8. What is the purpose of addressing this question at this time? |
| **III. Questions That Probe Assumptions** | 1. What are you assuming?<br>2. What is Karen assuming?<br>3. What could we assume instead?<br>4. You seem to be assuming ___. Do I understand you correctly?<br>5. All of your reasoning depends on the idea that ___ . Why have you based your reasoning on ___ rather than ___ ?<br>6. You seem to be assuming ___. How would you justify taking this for granted?<br>7. Is it always the case? Why do you think the assumption holds here? |
| **IV. Questions That Probe Information, Reasons, Evidence, and Causes** | 1. What would be an example?<br>2. How do you know?<br>3. What are your reasons for saying that?<br>4. Why did you say that?<br>5. What other information do we need to know before we can address this question?<br>6. Why do you think that is true?<br>7. Could you explain your reasons to us?<br>8. What led you to that belief?<br>9. Is this good evidence for believing that?<br>10. Do you have any evidence to support your assertion?<br>11. Are those reasons adequate?<br>12. How does that information apply to this case?<br>13. Is there reason to doubt that evidence?<br>14. What difference does that make?<br>15. Who is in a position to know if that is the case?<br>16. What would convince you otherwise?<br>17. What would you say to someone who said ___ ?<br>18. What accounts for ___ ?<br>19. What do you think is the cause?<br>20. How did this come about?<br>21. By what reasoning did you come to that conclusion?<br>22. How could we go about finding out whether that is true?<br>23. Can someone else give evidence to support that response? |
| **V. Questions about Viewpoints or Perspectives** | 1. You seem to be approaching this issue from ___ perspective. Why have you chosen this perspective rather than that perspective?<br>2. How would other groups or types of people respond? Why? What would influence them?<br>3. How could you answer the objection that ___ would make?<br>4. Can/did anyone see this another way?<br>5. What would someone who disagrees say?<br>6. What is an alternative?<br>7. How are Ken’s and Roxanne’s ideas alike? Different? |
| **VI. Questions That Probe Implications and Consequences** | 1. What are you implying by that?<br>2. When you say ___, are you implying ___?<br>3. But if that happened, what else would also happen as a result? Why?<br>4. What effect would that have?<br>5. Would that necessarily happen or only probably happen?<br>6. What is an alternative?<br>7. If this and this are the case, then what else must be true? |
| **VII. Questions about the Question** | 1. How can we find out?<br>2. Is this the same issue as ___?<br>3. How could someone settle this question?<br>4. Can we break this question down at all?<br>5. Is the question clear? Do we understand it?<br>6. How would ___ put the issue?<br>7. Is this question easy or difficult to answer? Why?<br>8. What does this question assume?<br>9. Would ___ put the question differently?<br>10. Why is this question important?<br>11. Does this question ask us to evaluate something?<br>12. Do we need facts to answer this?<br>13. Do we all agree that this is the question?<br>14. To answer this question, what other questions would we have to answer first?<br>15. I’m not sure I understand how you are interpreting the main question at issue. Could you explain your interpretation? |
| **VIII. Questions That Probe Concepts** | 1. What is the main idea we are dealing with?<br>2. Why/how is this idea important?<br>3. Do these two ideas conflict? If so, how?<br>4. What was the main idea guiding the thinking of the character in this story?<br>5. How is this idea guiding our thinking as we try to reason through this issue? Is this idea causing us problems?<br>6. What main theories do we need to consider in figuring out ___?<br>7. Are you using this term “___” in keeping with educated usage?<br>8. Which main distinctions should we draw in reasoning through this problem?<br>9. Which idea is this author using in her or his thinking? Is there a problem with it? |
| **IX. Questions That Probe Inferences and Interpretations** | 1. Which conclusions are we coming to about ___ ?<br>2. On what information are we basing this conclusion?<br>3. Is there a more logical inference we might make in this situation?<br>4. How are you interpreting her behavior? Is there another possible interpretation?<br>5. What do you think of ___?<br>6. How did you reach that conclusion?<br>7. Given all the facts, what is the best possible conclusion?<br>8. How shall we interpret these data? |

## PDF Parsing Tools
1. [Docling](https://github.com/DS4SD/docling) - IBM-backed document parsing framework that extracts structured content (headings, tables, figures, and layout) into LLM-friendly formats such as Markdown and JSON.

2. [Unstructured](https://github.com/Unstructured-IO/unstructured) - General-purpose document ingestion framework that converts PDFs and other document formats into semantic elements for RAG pipelines.

3. [MinerU](https://github.com/opendatalab/MinerU) - High-accuracy PDF parser designed for complex and academic documents with Markdown and structured JSON output.

4. [Marker](https://github.com/VikParuchuri/marker) - Converts PDFs into high-quality Markdown while preserving tables, equations, and document structure.

5. [PyMuPDF4LLM](https://github.com/pymupdf/PyMuPDF4LLM) - Lightweight PDF-to-Markdown extraction library optimized for LLM preprocessing workflows.

6. [PaddleOCR](https://github.com/PaddlePaddle/PaddleOCR) - OCR toolkit used to extract text from scanned PDFs and image-based documents.

7. [Surya OCR](https://github.com/VikParuchuri/surya) - Modern OCR and document layout analysis system optimized for AI document processing pipelines.

8. [GROBID](https://github.com/kermitt2/grobid) - Specialized parser for scientific papers that extracts metadata, sections, citations, and references.

9. [ParseMyPDF](https://github.com/genieincodebottle/parsemypdf) - Benchmarking framework that compares outputs from multiple PDF parsers on the same documents.

10. [PDFPlumber](https://github.com/jsvine/pdfplumber) - Python library focused on extracting text, tables, and positional information from PDFs.

11. [Apache Tika](https://github.com/apache/tika) - Content extraction toolkit that parses PDFs and many other document formats into plain text and metadata.

12. [pdfminer.six](https://github.com/pdfminer/pdfminer.six) - Low-level Python library for detailed PDF text extraction and layout analysis.

13. [LlamaParse](https://github.com/run-llama/llama_parse) - LLM-oriented document parser that converts PDFs into structured formats suitable for retrieval systems.

14. [Nougat](https://github.com/facebookresearch/nougat) - OCR model that converts scientific PDF pages directly into structured markup.

15. [LayoutParser](https://github.com/Layout-Parser/layout-parser) - Deep-learning-based document layout analysis toolkit for identifying document regions and structure.


## Uncategorised
* [Roadmap references - roadmap.sh](https://roadmap.sh/)
* [Typing practise - entertained.app](https://entertrained.app/books)

## Python
In virtual env, if locally installed pip not working then do
```
python -m ensurepip --upgrade
python -m pip install --upgrade pip
```

Add dependency in uv
```
uv add <package name>
eg. uv add sentence-transformers
```


Regex to convert date formats
```
import re
from datetime import datetime

def format_date(date_str: str) -> str:
    """Converts various date formats into the DD-MON-YYYY format."""
    date_str = date_str.strip()
    
    # 1. Format: Mon DD,YYYY
    if re.match(r'^[A-Za-z]{3}\s+\d{1,2},\d{4}$', date_str):
        date_obj = datetime.strptime(date_str.replace(',', ' '), '%b %d %Y')
    
    # 2. Format: DD MON YYYY
    elif re.match(r'^\d{1,2}\s+[A-Z]{3}\s+\d{4}$', date_str):
        date_obj = datetime.strptime(date_str, '%d %b %Y')
        
    # 3. Format: DDMMYY 
    elif re.match(r'^\d{6}$', date_str):
        date_obj = datetime.strptime(date_str, '%d%m%y')

    # 4. Format: DD Mon YYYY (e.g., '24 Jun 2025')
    elif re.match(r'^\d{1,2}\s+[A-Za-z]{3}\s+\d{4}$', date_str):
        date_obj = datetime.strptime(date_str, '%d %b %Y')
        
    else:
        return date_str 

    return date_obj.strftime('%d-%b-%Y').upper()
```
