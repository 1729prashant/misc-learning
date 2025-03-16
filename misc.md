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


## iOS, Swift

Description | Link
------------- | -------------
Collaborative List of Open-Source iOS Apps | <https://github.com/dkhamsing/open-source-ios-apps>
21 Amazing Open Source iOS Apps Written in Swift | <https://medium.mybridge.co/21-amazing-open-source-ios-apps-written-in-swift-5e835afee98e>
App Icon Generator | <https://www.appicon.co/#app-icon>
Human Interface Guidelines, Apple | <https://developer.apple.com/design/human-interface-guidelines>
Swift map(_:), Apple | <https://developer.apple.com/documentation/applearchive/archiveheader/4339557-map/>
Swift Closures, Apple | <https://docs.swift.org/swift-book/documentation/the-swift-programming-language/closures/>
SwiftUI Tutorials, Apple| <https://developer.apple.com/tutorials/SwiftUI#//apple_ref/doc/uid/TP40015214-CH6-SW1>
Swift Thread, Apple | <https://developer.apple.com/documentation/foundation/thread>
Mac Catalyst, Apple | <https://developer.apple.com/documentation/uikit/mac_catalyst>
Swift ScrollView, Apple | <https://developer.apple.com/documentation/swiftui/scrollview>
Swift List, Apple | <https://developer.apple.com/documentation/swiftui/list>
Swift Navigation | <https://developer.apple.com/documentation/swiftui/navigation>
Swift Gestures, Apple | <https://developer.apple.com/documentation/swiftui/gestures>
Swift CaseIterable, Apple | <https://developer.apple.com/documentation/swift/caseiterable>
Swift Type Casting | <https://docs.swift.org/swift-book/documentation/the-swift-programming-language/typecasting/>
Swift Access Control | <https://docs.swift.org/swift-book/documentation/the-swift-programming-language/accesscontrol/>
Swift Protocols | <https://docs.swift.org/swift-book/documentation/the-swift-programming-language/protocols>
Swift Error Handling | <https://docs.swift.org/swift-book/documentation/the-swift-programming-language/errorhandling/>
Swift Performance Tests, Apple | <https://developer.apple.com/documentation/xctest/performance_tests>
Swift User Interface TEsts, Apple | <https://developer.apple.com/documentation/xctest/user_interface_tests>
Swift XCTest, Apple | <https://developer.apple.com/documentation/xctest>
Swift URL Loading System, Apple | <https://developer.apple.com/documentation/foundation/url_loading_system>
Swift Handling an authentication challenge, Apple | <https://developer.apple.com/documentation/foundation/url_loading_system/handling_an_authentication_challenge>
Swift Preventing Insecure Network Connections, Apple | <https://developer.apple.com/documentation/security/preventing-insecure-network-connections>
Swift URLRequest, Apple | <https://developer.apple.com/documentation/foundation/urlrequest>
Swift Fetching website data into memory, Apple | <https://developer.apple.com/documentation/foundation/url_loading_system/fetching_website_data_into_memory>
Swift URLSession, Apple | <https://developer.apple.com/documentation/foundation/urlsession>
Swift URLSessionTask, Apple | <https://developer.apple.com/documentation/foundation/urlsessiontask>
Swift Core Data, Apple | <https://developer.apple.com/documentation/coredata>
Swift SwiftData, Apple | <https://developer.apple.com/documentation/SwiftData>
Swift Configuring Entities, Apple| <https://developer.apple.com/documentation/coredata/modeling_data/configuring_entities>
Swift Configuring Attributes, Apple | <https://developer.apple.com/documentation/coredata/modeling_data/configuring_attributes>
Swift Configuring Relationships, Apple | <https://developer.apple.com/documentation/coredata/modeling_data/configuring_relationships>
Swift NSManagedObjectContext, Apple | <https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext>
Creating Predicates, Apple | <https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/Articles/pCreating.html>
Using Predicates, Apple | <https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/Articles/pUsing.html>
Swift NSComparisonPredicate, Apple | <https://developer.apple.com/documentation/foundation/nscomparisonpredicate>
Swift SortDescriptor, Apple | <https://developer.apple.com/documentation/foundation/sortdescriptor>
Submit your iOS apps to the App Store, Apple | <https://developer.apple.com/ios/submit/>
App Review Guidelines, Apple | <https://developer.apple.com/app-store/review/guidelines/>
App Store Connect Help, Apple | <https://developer.apple.com/help/app-store-connect/>
TestFlight, Apple | <https://developer.apple.com/testflight/>
Swift Forums | <https://forums.swift.org/top?period=weekly>
Xcode Releases | <https://xcodereleases.com/>
Terminal Cheatsheet for Mac (Basics) | <https://github.com/0nn0/terminal-mac-cheatsheet>
Side Projects, reddit | <https://www.reddit.com/r/SideProject/>
Flat UI Colors | <https://flatuicolors.com/>
Google Fonts | <https://fonts.google.com/>
CS193p - Developing Apps for iOS, Stanford | <https://cs193p.sites.stanford.edu/>
Nil Coalescing | <https://nilcoalescing.com/>
List of free and open-source iOS applications, Wikipedia | <https://en.wikipedia.org/wiki/List_of_free_and_open-source_iOS_applications>
Alamofire is an HTTP networking library written in Swift. | <https://github.com/Alamofire/Alamofire>


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


## Roadmaps
* <https://roadmap.sh/>

