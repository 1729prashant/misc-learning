
## PeopleSoft

#### Tables
```
Projects
---------
PSPROJECTDEFN — Project header table
PSPROJECTITEM — Definitions in the project 
PS_PTOBJECTTYPES — Object types list

Fields
---------
PSDBFIELD — Fields in the system
PSDBFLDLABL — Field Labels
PSDBFIELDLANG
PSXLATITEM — Translate Values
PSXLATITEMLANG
PSDBFLDLABLLANG

Records
---------
PSRECDEFN — Record header table (record types)
PSRECFIELD — Fields in the record (subrecords not expanded)
PSRECFIELDALL — Fields in the record (subrecords expanded)
PSKEYDEFN — Indexes
PSTBLSPCCAT — Tablespaces
PSRECTBLSPC — Records’ tablespace assignments
PSRECDEFNLANG — language Record

Pages
---------
PSPNLDEFN — Page header table
PSPNLFIELD — Page controls (field types/FIELDTYPE)
PSPNLHTMLAREA — Static HTML Areas on Pages
PSPNLCNTRLDATA — stores additional attributes for Page controls
PSPNLFIELDEXT — stores additional attributes for Page controls
PSPTPNLRTEDITOR — stores attributes relating to the rich-text editor for long character fields.

Components
---------
PSPNLGRPDEFN — Component header table
PSPNLGROUP — Pages in the components

Component Interface
---------
PSBCDEFN — header record; one row for each component interface
PSBCITEM — one row for each property

Menus
---------
PSMENUDEFN — Menu header table
PSMENUITEM — Items (components) on the menu

Security
---------
PSCLASSDEFN — Permission List header table
PSAUTHITEM — Menu items granted security by permission lists
PSROLEDEFN — Role header table
PSROLECLASS — Permission Lists in roles
PSOPRDEFN — User ID header table
PSROLEUSER — Roles granted to users
PSAUTHBUSCOMP — Access to Component Interfaces
PSAUTHWEBLIBVW — Access to just Web Libraries
PS_SCRTY_ACC_GRP — Query Security Access Groups
PSOPROBJ — Definition security (connects a permission list to a definition group)


PeopleSoft Login Details and Authorization
---------
PSACCESSLOG	- Login and logout information of the users
PSACCESSPRFL - Contains the symbolic id,accessid/password details.
PSCLOCK	- Login fails if the table is empty
PSOPRDEFN	- This table holds the peoplesoft Oprid’s/passwords info with symbolic id.


Process Scheduler
---------
PS_PRCSDEFN — Process Definition Header
PS_PRCSTYPEDEFN — Process Type Definition Header
PS_PRCSDEFNGRP — Process Group
PS_PRCSDEFNPNL — Component
PS_PRCSJOBDEFN — Job Header
PSPRCSRQST — Process Request Instances
PS_PRCSJOBITEM — Job Processes
PSREN — stores REN server ports
PS_SERVERCLASS — Server Class
PS_PRCS_CAT_TBL — Process Categories
PS_SERVERCATEGORY — Server Categories
PS_SERVERDEFN — Server Definition Header
PSPRCSQUE — Process Queue
PSRF_FLIST_TBL — Folder Definition

Portal
---------
PSPRSMDEFN — Content References and Folders
PSPRUHTABPGLT — Portal User HP Tab Pagelet
PSPRUHDEFN — Homepage definition (from here)
PSPRUHTAB — Homepage Tab (from here)
PSWEBPROFNVP — Web Profile Settings

Change Control
---------
PSCHGCTLHIST — shows history of locked definitions with project name, incident, and description
PSCHGCTLLOCK — shows definitions that are currently locked
PSCOMPOBJDIFF — I think this shows compare report from a Project Compare
PSPROJECTMSG — messages from a project copy

Application Engine
---------
PSAEAPPLDEFN — header record; 1 row per app engine
PSAEAPPLSTATE — state records assigned to app engines
PSAEAPPLTEMPTBL — temp tables assigned to app engines
PSAESECTDEFN — sections
PSAESTEPDEFN — steps
PSAESTEPMSGDEFN
PSAESTMTDEFN — actions (action types)
PSAESECTDTLDEFN — section details
PS_AEONLINEINST 
PS_AEINSTANCENBR 
PS_AELOCKMGR 

HTML Definitions
---------
PSCONTDEFN — header record; last update time, etc.
PSCONTENT — stores actual text in the HTML definition

PeopleCode
---------
PSPCMPROG — PeopleCode program
PSPCMTXT — The plain text of the PeopleCode programs

SQL Definitions
---------
PSSQLDEFN — header record; last update time, etc.
PSSQLTEXTDEFN — stores actual text in the SQL definition
PS_SQLSTMT_TBL — stores SQL statemernts for COBOL

File Layout Definitions
---------
PSFLDDEFN — header record; last update time, etc.
PSFLDSEGDEFN — stores the segments for each layout
PSFLDFIELDDEFN — stores the fields for each layout

Query
---------
PSQRYDEFN — Query Definition Header
PSQRYBIND — Query bind variables/parameters
PSQRYRECORD — Query records
PSQRYFIELD — Fields
PSCONQRSDEFN — connected query
PSCONQRSMAP — connected query: list of child queries
PSCONQRSFLDREL — connected query
PSCONQRSRUNCNTR — connected query
PSCONQRSRUNPRM — connected query

Workflow
---------
APPR_RULE_DETL – Approval Rule Defn Details
APPR_RULE_FIELD – Approval Rule Defn Route Cntl
APPR_RULE_AMT – Approval Rule Amounts
RTE_CNTL_LN – Route Control Profile Line
RTE_CNTL_RUSER – RoleUser Route Cntl Profiles
RTE_CNTL_TYPE – Route Control Type
RTE_CNTL_HDR – Routing Control Type
PSWORKLIST — list of work items for each user
PS_WF_INSTSTATUS — description of the status
PSACTIVITYDEFN — Activity Definition
PSBUSPROCDEFN — Business Process Definition
PSEVENTDEFN — Event Definition
PSEVENTROUTE — Event Route

Timings
---------
BAT_TIMINGS_LOG
BAT_TIMINGS_DTL
BAT_TIMINGS_FN
File References
PSFILEREDEFN — File references (keys: FILEREFNAME, FILEREFTYPECODE)
EDI Manager
PS_ECINMAPFILE — Records
PS_ECINMAPREC — Records
PS_ECMAPDEFN — EDI Map Definition
PS_ECINMAPFLD — Fields in the Records
PS_ECINMAPRECFLD — Fields in the Records
PS_ECOUTMAPREC
PS_ECOUTMAPFLD

Feeds
---------
PS_PTFP_FEED
PS_PTFP_DATATYPE
PS_PTFP_DTYPE_IBSO
PS_PTFP_DTYPE_ATTR
PS_PTFP_ATTRS
PS_PTFP_SETTINGS
PS_PTFP_PARMS
PS_PTFP_SECURITY
PS_PTFP_PVALS
PS_PTFP_ADMN_PREF
PS_PTFP_USER_PREF

Integration Broker
---------
PSMSGPARTS — Message Parts
PSMSGVER — Message Version
PSMSGREC — Message Record
PSOPERATION — Operation
PSOPRVERDFN — Operation Versions
PSSERVICEOPR — Service Operations
PSOPRHDLR — Operation Handlers
PSIBRTNGDEFN — Routings
PSMSGNODEDEFN — Message Nodes
PSRTNGDFNPARM — Routing Parameters
PSOPRVERDFNPARM
PSQUEUEDEFN
PSIBSCMADFN
PSOPERATIONDMS

Message Catalog
---------
PSMSGSETDEFN — Message Set — the header for that message set
PSMSGCATDEFN — Individual messages in the message catalog/message set

SQR Strings
---------
PS_STRINGS_TBL — String definition  (sqrtarns.sqc)
PS_STRINGS_LNG_TBL — Translated String  (sqrtarns.sqc)

Sources:
---------
for PS tables: see mvprdexp.dms in the home scripts directory
```

#### SQL
```
sql for effective active translate values
-----------------------------------------
select
A.*
from PSXLATITEM A
where EFFDT = (
      select max(EFFDT)
      from PSXLATITEM
      where FIELDNAME = A.FIELDNAME
      and FIELDVALUE = A.FIELDVALUE
      and A.EFFDT <= sysdate
      )
and A.EFF_STATUS = 'A';


select
A.*
from
PSXLATITEM A, PSRECFIELDALL B
where A.FIELDNAME=b.FIELDNAME
      and B.recname= 'PAYMENT_TBL'
      and A.EFFDT = (
          select max(EFF.EFFDT)
          from PSXLATITEM EFF
          where EFF.FIELDNAME = A.FIELDNAME
          and EFF.FIELDVALUE = A.FIELDVALUE
          and A.EFFDT <= sysdate
          )


--SQL for FIELD PROPERTIES
---------------------------
SELECT DISTINCT
    FD.FIELDNAME,
    CASE FD.FIELDTYPE
        WHEN 0
            THEN 'Char'
        WHEN 1
            THEN 'Long Char'
        WHEN 2
            THEN 'NBR'
        WHEN 3
            THEN 'Sign'
        WHEN 4
            THEN 'Date'
        WHEN 5
            THEN 'Time'
        WHEN 6
            THEN 'Dttm'
        WHEN 7
            THEN 'Img'
        WHEN 8
            THEN 'Vers'
        ELSE
            ' '
    END 'TYPE',
    FD.LENGTH,
    FL.SHORTNAME,
    FL.LONGNAME
FROM
    PSDBFIELD FD,
    PSDBFLDLABL FL,
    PSPROJECTITEM PI
WHERE
    FD.FIELDNAME = FL.FIELDNAME
    AND PI.OBJECTVALUE1 = FD.FIELDNAME
    AND PI.OBJECTTYPE = 2
    AND PI.OBJECTID1 = 6
    AND FD.FIELDNAME = ' ' --fieldname here
ORDER BY
    FD.FIELDNAME;

--SQL for TRANSLATE VALUES
---------------------------
SELECT DISTINCT
    X.FIELDVALUE,
    X.EFFDT,
    X.EFF_STATUS,
    X.XLATSHORTNAME,
    X.XLATLONGNAME
FROM
    PSPROJECTITEM A,
    PSXLATITEM X
WHERE
    A.OBJECTTYPE = 4
    AND A.OBJECTID1 = 6
    AND X.EFFDT =
        (
            SELECT
                MAX(B_ED.EFFDT)
            FROM
                PSXLATITEM B_ED
            WHERE
                X.FIELDNAME = B_ED.FIELDNAME
                AND X.FIELDVALUE = B_ED.FIELDVALUE
                AND B_ED.EFFDT <= SYSDATE
        )
    AND X.FIELDNAME = ' '; --here

--SQL for RECORD FIELD PROPERTIES
------------------------------------
SELECT DISTINCT
    PR.RECNAME,
    PR.FIELDNUM,
    PR.FIELDNAME,
    CASE MOD(PR.USEEDIT, 2)
        WHEN 1
            THEN 'Y'
        ELSE
            ' '
    END 'KEY/REQ',
    CASE MOD((PR.USEEDIT / 256), 2)
        WHEN 1
            THEN 'R'
        ELSE
            ' '
    END 'KEY/REQ',
    CASE MOD((PR.USEEDIT / 2048), 2)
        WHEN 1
            THEN 'YES'
        ELSE
            ' '
    END 'SEARCH KEY',
    CASE MOD((PR.USEEDIT / 32), 2)
        WHEN 1
            THEN 'YES'
        ELSE
            ' '
    END 'LIST BOX ITEM',
    CASE MOD((PR.USEEDIT / 16), 2)
        WHEN 1
            THEN 'ALT'
        ELSE
            ' '
    END 'ALT SEARCH KEY'
FROM
    PSRECFIELD PR,
    PSPROJECTITEM A,
    PSRECDEFN D
WHERE
    PR.RECNAME = D.RECNAME
    AND A.OBJECTTYPE = 0
    AND A.OBJECTID1 = 1
    AND PR.RECNAME = '' -- record name here
ORDER BY
    PR.RECNAME,
    PR.FIELDNUM;

--PAGE FIELD PROPERTIES 
---------------------------
SELECT distinct
    PF.PNLNAME,
    PF.FIELDNUM,
    PF.OCCURSLEVEL AS 'LVL',
    PF.LBLTEXT     AS 'LABEL',
    CASE PF.FIELDTYPE
        WHEN 1
            THEN 'Frame'
        WHEN 2
            THEN 'Group Box'
        WHEN 3
            THEN 'Static Image'
        WHEN 4
            THEN 'Edit Box'
        WHEN 5
            THEN 'Drop Down List'
        WHEN 6
            THEN 'Long Edit Box'
        WHEN 7
            THEN 'Check Box'
        WHEN 8
            THEN 'Radio Button'
        WHEN 9
            THEN 'Image'
        WHEN 10
            THEN 'Scroll Bar'
        WHEN 11
            THEN 'SubPage'
        WHEN 12
            THEN 'Push Button/Hyper Link'
        WHEN 13
            THEN 'Push Button/Hyper Link'
        WHEN 14
            THEN 'Push Button/Hyper Link'
        WHEN 15
            THEN 'Push Button/Hyper Link'
        WHEN 16
            THEN 'Push Button/Hyper Link'
        WHEN 17
            THEN 'Text'
        WHEN 18
            THEN 'SecPage'
        WHEN 19
            THEN 'Grid'
        WHEN 20
            THEN 'Tree'
        WHEN 21
            THEN 'Push Button/HyperLink'
        WHEN 23
            THEN 'Horizontle Rule'
        WHEN 24
            THEN 'Tab Separator'
        WHEN 25
            THEN 'HTML Area'
        WHEN 26
            THEN 'Push Button/Hyper Link'
        WHEN 27
            THEN 'Scroll Area'
        ELSE
            ' '
    END            'TYPE',
    PF.FIELDNAME,
    PF.RECNAME
FROM
    PSPNLFIELD PF,
    PSPROJECTITEM A,
    PSPNLDEFN F
WHERE
    F.PNLNAME = PF.PNLNAME
    AND A.OBJECTTYPE = 5
    AND A.OBJECTID1 = 9
    AND F.PNLNAME = '' --here
ORDER BY
    PF.PNLNAME,
    PF.FIELDNUM;

--SQL for COMPONENT PROPERTIES
-------------------------------
SELECT DISTINCT
    MI.PNLGRPNAME     AS 'COMPONENT',
    MI.ITEMLABEL      AS 'DESCRIPTION',
    PG.ITEMLABEL      AS 'ITEM LABEL',
    PG.FOLDERTABLABEL AS 'FOLDER TAB LABEL',
    H.SEARCHRECNAME   AS 'SEARCH RECORD',
    PG.PNLNAME        AS 'PAGE NAME'
FROM
    PSMENUITEM MI,
    PSPNLGROUP PG,
    PSPNLGRPDEFN H,
    PSPROJECTITEM A
WHERE
    MI.PNLGRPNAME = PG.PNLGRPNAME
    AND PG.PNLGRPNAME = H.PNLGRPNAME
    AND PG.MARKET = H.MARKET
    AND A.OBJECTTYPE = 7
    AND A.OBJECTID1 = 10
    AND PG.PNLGRPNAME = ''; --component name here

--SQL for FIELDS FROM A PROJECT
------------------------------------
select
    projectname,
    'Field',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 2
    and objectid1 = 6
    and projectname = 'project name here';

-- fields in project
select
    projectname,
    'Field Translate Values',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 4
    and objectid1 = 6
    and projectname = 'project name here';

--RECORDS IN PROJECT
select
    projectname,
    'Record',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 0
    and objectid1 = 1
    and projectname = 'project name here';

--RECORD VIEW SQL IN PROJECT
select
    projectname,
    'Record View SQL',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 30
    and objectid1 = 65
    and projectname = ' project name here';

--PAGE IN PROJECT
select
    projectname,
    'Page',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 5
    and objectid1 = 9
    and projectname = ' project name here';

--COMPONENT IN PROJECT
select
    projectname,
    'Component',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 7
    and objectid1 = 10
    and projectname = ' project name here';

--MENU IN PROJECT
select
    projectname,
    'Menu',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 6
    and objectid1 = 3
    and projectname = ' project name here';

--RECORD FIELD PEOPLECODE
select
    projectname,
    'Record Field Peoplecode',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 8
    and objectid1 = 1
    and projectname = ' project name here'
select
    projectname,
    'Record Field Peoplecode',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 44
    and objectid1 = 9
    and projectname = ' project name here';

--QUERY
select
    projectname,
    'Query',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 10
    and objectid1 = 30
    and projectname = ' project name here';

--COMPONENT INTERFACE
select
    projectname,
    'Component Interface',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 32
    and objectid1 = 74
    and projectname = ' project name here';

--MESSAGE CATALOG ENTRIES
select
    projectname,
    'Message Catalog Entries',
    objectvalue1,
    Objectvalue2,
    objectvalue3
from
    psprojectitem
where
    objecttype = 25
    and objectid1 = 48
    and projectname = 'project name here';


--SQL for PAGE ACCESS 
-----------------------
/**
Input : -User ID
Page Name for which you want to check the access type.
Notations used in the query output.
-------------
C Correction
U-DA Update Display All
UD Update Display
A Add
-------------
**/

SELECT
    A.ROLEUSER,
    C.PNLITEMNAME,
    A.ROLENAME,
    B.CLASSID,
    CASE
        WHEN C.AUTHORIZEDACTIONS = 15
            THEN 'C|U-DA|UD|A'
        ELSE
            CASE
                WHEN C.AUTHORIZEDACTIONS = 14
                    THEN 'C|U-DA|UD| '
                ELSE
                    CASE
                        WHEN C.AUTHORIZEDACTIONS = 13
                            THEN 'C|U-DA| |A'
                        ELSE
                            CASE
                                WHEN C.AUTHORIZEDACTIONS = 12
                                    THEN 'C|U-DA| | '
                                ELSE
                                    CASE
                                        WHEN C.AUTHORIZEDACTIONS = 11
                                            THEN 'C| |UD|A'
                                        ELSE
                                            CASE
                                                WHEN C.AUTHORIZEDACTIONS = 10
                                                    THEN 'C| |UD| '
                                                ELSE
                                                    CASE
                                                        WHEN C.AUTHORIZEDACTIONS = 9
                                                            THEN 'C| | |A'
                                                        ELSE
                                                            CASE
                                                                WHEN C.AUTHORIZEDACTIONS = 8
                                                                    THEN 'C| | | '
                                                                ELSE
                                                                    CASE
                                                                        WHEN C.AUTHORIZEDACTIONS = 7
                                                                            THEN '|U-DA|UD|A'
                                                                        ELSE
                                                                            CASE
                                                                                WHEN C.AUTHORIZEDACTIONS = 6
                                                                                    THEN ' |U-DA|UD| '
                                                                                ELSE
                                                                                    CASE
                                                                                        WHEN C.AUTHORIZEDACTIONS = 5
                                                                                            THEN ' |U-DA| |A'
                                                                                        ELSE
                                                                                            CASE
                                                                                                WHEN C.AUTHORIZEDACTIONS = 4
                                                                                                    THEN ' |U-DA| | '
                                                                                                ELSE
                                                                                                    CASE
                                                                                                        WHEN C.AUTHORIZEDACTIONS = 3
                                                                                                            THEN ' | |UD|A'
                                                                                                        ELSE
                                                                                                            CASE
                                                                                                                WHEN C.AUTHORIZEDACTIONS = 2
                                                                                                                    THEN ' | |UD| '
                                                                                                                ELSE
                                                                                                                    CASE
                                                                                                                        WHEN C.AUTHORIZEDACTIONS = 1
                                                                                                                            THEN ' | | |A'
                                                                                                                        ELSE
                                                                                                                            CASE
                                                                                                                                WHEN C.AUTHORIZEDACTIONS = 0
                                                                                                                                    THEN ' | | | '
                                                                                                                                ELSE
                                                                                                                                    'Invalid AUTHORIZATION Type'
                                                                                                                            END
                                                                                                                    END
                                                                                                            END
                                                                                                    END
                                                                                            END
                                                                                    END
                                                                            END
                                                                    END
                                                            END
                                                    END
                                            END
                                    END
                            END
                    END
            END
    END AS ALLOWED
FROM
    PSROLEUSER A,
    PSROLECLASS B,
    PSAUTHITEM C
WHERE
    C.CLASSID = B.CLASSID
    AND B.ROLENAME = A.ROLENAME
    AND C.PNLITEMNAME = ' ' --component name here
    AND A.ROLEUSER = 'oprid goes here';

--SQL to get the Key for a particular table
-------------------------------------------
select
    *
from
    pskeydefn
where
    recname in (
                   'PSACTIVITYDEFN', 'PSACTIVITYLANG', 'PSAEAPPLDEFN', 'PSAEAPPLLANG', 'PSAESECTDTLDEFN',
                   'PSAESTEPDEFN', 'PSAESTMTDEFN', 'PSBCDEFN', 'PSBCDEFNLANG', 'PSBUSPITEMLANG', 'PSBUSPROCDEFN',
                   'PSBUSPROCITEM', 'PSBUSPROCLANG', 'PSCHGCTLHIST', 'PSCHGCTLLOCK', 'PSCHNLDEFN', 'PSCHNLDEFNLANG',
                   'PSDBFIELD', 'PSDBFIELDLANG', 'PSFLDDEFN', 'PSMENUDEFN', 'PSMENUDEFNLANG', 'PSMSGDEFN',
                   'PSMSGDEFNLANG', 'PSMSGNODEDEFN', 'PSNODEDEFNLANG', 'PSOPRALIASTYPE', 'PSOPRALSTYPLANG',
                   'PSPNLDEFN', 'PSPNLGDEFNLANG', 'PSPNLGRPDEFN', 'PSPROJDEFNLANG', 'PSPROJECTDEFN', 'PSQRYDEFN',
                   'PSQRYDEFNLANG', 'PSRECDEFN', 'PSRECDEFNLANG', 'PSROLEDEFN', 'PSROLEDEFN', 'PSSQLDESCR',
                   'PSSQLLANG', 'PSVAITEM', 'PSVAITEMLANG'
               )
order by
    keyposn;

--SQL for all the Fields in a Project
-------------------------------------
select
    A.projectname,
    B.fieldname
from
    psprojectitem A,
    psdbfield B
where
    A.objectvalue1 = B.fieldname
    and A.objecttype = 2
    and A.objectid1 = 6;

--SQL for message catalog
-------------------------
SELECT
    *
FROM
    PSMSGSETDEFN;
SELECT
    MESSAGE_TEXT,
    LANGUAGE_CD
FROM
    PSMSGCATLANG
WHERE
    MESSAGE_SET_NBR = 22001
    AND MESSAGE_NBR = 1;
SELECT
    MESSAGE_TEXT
FROM
    PSMSGCATDEFN
WHERE
    MESSAGE_SET_NBR = 22001
    AND MESSAGE_NBR = 1;

--SQL for required fields
-------------------------
SELECT
    FIELDNAME,
    CASE MOD((USEEDIT / 256), 2)
        WHEN 1
            THEN 'Y'
        ELSE
            'N'
    END as REQUIRED
FROM
    PSRECFIELDDB
WHERE
    RECNAME = 'VOUCHER';


TBD...

```



## Oracle SQL Fuzzy search

use UTL_MATCH
https://docs.oracle.com/en/database/oracle/oracle-database/12.2/arpls/UTL_MATCH.html#GUID-98BCC5BB-5A0A-422B-99FD-3D85F282923A

use FUZZY_MATCH (better as of 2024)
https://docs.oracle.com/en/database/oracle/oracle-database/23/sqlrf/data-quality-operators.html#GUID-C13A179C-1F82-4522-98AA-E21C6504755E



## SQLite Limitations

* <https://www.sqlite.org/omitted.html> - SQL Features That SQLite Does Not Implement
* <https://www.sqlite.org/quirks.html>
