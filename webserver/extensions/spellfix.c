<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<base href="https://www.sqlite.org/src/doc/trunk/README.md">
<meta http-equiv="Content-Security-Policy" content="default-src 'self' data:; script-src 'self' 'nonce-5c468789942c60a9e666383658ec74ab1612bc120dd0b847'; style-src 'self' 'unsafe-inline'; img-src * data:">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>SQLite: Documentation</title>
<link rel="alternate" type="application/rss+xml" title="RSS Feed"  href="/src/timeline.rss">
<link rel="stylesheet" href="/src/style.css?id=cbbd4405" type="text/css">
</head>
<body class="doc rpage-doc cpage-doc">
<header>
  <div class="title"><h1>SQLite</h1>Documentation</div>
  <div class="status">
    <a href='/src/login'>Login</a>

  </div>
</header>
<nav class="mainmenu" title="Main Menu">
  <a id='hbbtn' href='/src/sitemap' aria-label='Site Map'>&#9776;</a><a href='https://sqlite.org/' class=''>Home</a>
<a href='/src/timeline' class=''>Timeline</a>
<a href='https://sqlite.org/forum/forum' class='desktoponly'>Forum</a>

</nav>
<nav id="hbdrop" class='hbdrop' title="sitemap"></nav>
<div class="content"><span id="debugMsg"></span>
<div class="markdown">

<p><h1 align="center">SQLite Source Repository</h1></p>

<p>This repository contains the complete source code for the
<a href="https://sqlite.org/">SQLite database engine</a>, including
many test scripts.  However, other test scripts
and most of the documentation are managed separately.</p>

<p>See the <a href="https://sqlite.org/">on-line documentation</a> for more information
about what SQLite is and how it works from a user's perspective.  This
README file is about the source code that goes into building SQLite,
not about how SQLite is used.</p>

<h2>Version Control</h2>
<p>SQLite sources are managed using
<a href="https://fossil-scm.org/">Fossil</a>, a distributed version control system
that was specifically designed and written to support SQLite development.
The <a href="https://sqlite.org/src/timeline">Fossil repository</a> contains the urtext.</p>

<p>If you are reading this on GitHub or some other Git repository or service,
then you are looking at a mirror.  The names of check-ins and
other artifacts in a Git mirror are different from the official
names for those objects.  The official names for check-ins are
found in a footer on the check-in comment for authorized mirrors.
The official check-in name can also be seen in the <code>manifest.uuid</code> file
in the root of the tree.  Always use the official name, not  the
Git-name, when communicating about an SQLite check-in.</p>

<p>If you pulled your SQLite source code from a secondary source and want to
verify its integrity, there are hints on how to do that in the
<a href="#vauth">Verifying Code Authenticity</a> section below.</p>

<h2>Contacting The SQLite Developers</h2>
<p>The preferred way to ask questions or make comments about SQLite or to
report bugs against SQLite is to visit the 
<a href="https://sqlite.org/forum">SQLite Forum</a> at <a href="https://sqlite.org/forum/">https://sqlite.org/forum/</a>.
Anonymous postings are permitted.</p>

<p>If you think you have found a bug that has security implications and
you do not want to report it on the public forum, you can send a private
email to drh at sqlite dot org.</p>

<h2>Public Domain</h2>
<p>The SQLite source code is in the public domain.  See
<a href="https://sqlite.org/copyright.html">https://sqlite.org/copyright.html</a> for details. </p>

<p>Because SQLite is in the public domain, we do not normally accept pull
requests, because if we did take a pull request, the changes in that
pull request might carry a copyright and the SQLite source code would
then no longer be fully in the public domain.</p>

<h2>Obtaining The SQLite Source Code</h2>
<p>If you do not want to use Fossil, you can download tarballs or ZIP
archives or <a href="https://sqlite.org/cli.html#sqlar">SQLite archives</a> as follows:</p>

<ul>
<li><p>Latest trunk check-in as
 <a href="https://www.sqlite.org/src/tarball/sqlite.tar.gz">Tarball</a>,
 <a href="https://www.sqlite.org/src/zip/sqlite.zip">ZIP-archive</a>, or
 <a href="https://www.sqlite.org/src/sqlar/sqlite.sqlar">SQLite-archive</a>.</p></li>
<li><p>Latest release as
 <a href="https://www.sqlite.org/src/tarball/sqlite.tar.gz?r=release">Tarball</a>,
 <a href="https://www.sqlite.org/src/zip/sqlite.zip?r=release">ZIP-archive</a>, or
 <a href="https://www.sqlite.org/src/sqlar/sqlite.sqlar?r=release">SQLite-archive</a>.</p></li>
<li><p>For other check-ins, substitute an appropriate branch name or
 tag or hash prefix in place of "release" in the URLs of the previous
 bullet.  Or browse the <a href="https://www.sqlite.org/src/timeline">timeline</a>
 to locate the check-in desired, click on its information page link,
 then click on the "Tarball" or "ZIP Archive" links on the information
 page.</p></li>
</ul>

<p>To access sources directly using <a href="https://fossil-scm.org/home">Fossil</a>,
first install Fossil version 2.0 or later.
Source tarballs and precompiled binaries available at
<a href="https://fossil-scm.org/home/uv/download.html">https://fossil-scm.org/home/uv/download.html</a>.  Fossil is
a stand-alone program.  To install, simply download or build the single
executable file and put that file someplace on your $PATH.
Then run commands like this:</p>

<pre><code>    mkdir -p ~/sqlite
    cd ~/sqlite
    fossil open https://sqlite.org/src
</code></pre>

<p>The "fossil open" command will take two or three minutes.  Afterwards,
you can do fast, bandwidth-efficient updates to the whatever versions
of SQLite you like.  Some examples:</p>

<pre><code>    fossil update trunk             ;# latest trunk check-in
    fossil update release           ;# latest official release
    fossil update trunk:2024-01-01  ;# First trunk check-in after 2024-01-01
    fossil update version-3.39.0    ;# Version 3.39.0
</code></pre>

<p>Or type "fossil ui" to get a web-based user interface.</p>

<h2>Compiling for Unix-like systems</h2>
<p>First create a directory in which to place
the build products.  It is recommended, but not required, that the
build directory be separate from the source directory.  Cd into the
build directory and then from the build directory run the configure
script found at the root of the source tree.  Then run "make".</p>

<p>For example:</p>

<pre><code>    apt install gcc make tcl-dev  ;#  Make sure you have all the necessary build tools
    tar xzf sqlite.tar.gz         ;#  Unpack the source tree into "sqlite"
    mkdir bld                     ;#  Build will occur in a sibling directory
    cd bld                        ;#  Change to the build directory
    ../sqlite/configure           ;#  Run the configure script
    make sqlite3                  ;#  Builds the "sqlite3" command-line tool
    make sqlite3.c                ;#  Build the "amalgamation" source file
    make sqldiff                  ;#  Builds the "sqldiff" command-line tool
    # Makefile targets below this point require tcl-dev
    make tclextension-install     ;#  Build and install the SQLite TCL extension
    make devtest                  ;#  Run development tests
    make releasetest              ;#  Run full release tests
    make sqlite3_analyzer         ;#  Builds the "sqlite3_analyzer" tool
</code></pre>

<p>See the makefile for additional targets.  For debugging builds, the
core developers typically run "configure" with options like this:</p>

<pre><code>    ../sqlite/configure --enable-all --enable-debug CFLAGS='-O0 -g'
</code></pre>

<p>For release builds, the core developers usually do:</p>

<pre><code>    ../sqlite/configure --enable-all
</code></pre>

<p>Almost all makefile targets require a "tclsh" TCL interpreter version 8.6 or
later.  The "tclextension-install" target and the test targets that follow
all require TCL development libraries too.  ("apt install tcl-dev").  It is
helpful, but is not required, to install the SQLite TCL extension (the
"tclextension-install" target) prior to running tests.  The "releasetest"
target has additional requiremenst, such as "valgrind".</p>

<p>On "make" command-lines, one can add "OPTIONS=..." to specify additional
compile-time options over and above those set by ./configure.  For example,
to compile with the SQLITE_OMIT_DEPRECATED compile-time option, one could say:</p>

<pre><code>    ./configure --enable-all
    make OPTIONS=-DSQLITE_OMIT_DEPRECATED sqlite3
</code></pre>

<p>The configure script uses autoconf 2.61 and libtool.  If the configure
script does not work out for you, there is a generic makefile named
"Makefile.linux-gcc" in the top directory of the source tree that you
can copy and edit to suit your needs.  Comments on the generic makefile
show what changes are needed.</p>

<h2>Compiling for Windows Using MSVC</h2>
<p>On Windows, everything can be compiled with MSVC.
You will also need a working installation of TCL.
See the <a href="doc/compile-for-windows.md">compile-for-windows.md</a> document for
additional information about how to install MSVC and TCL and configure your
build environment.</p>

<p>If you want to run tests, you need to let SQLite know the location of your
TCL library, using a command like this:</p>

<pre><code>    set TCLDIR=c:\Tcl
</code></pre>

<p>SQLite uses "tclsh.exe" as part of the build process, and so that
program will need to be somewhere on your %PATH%.  SQLite itself
does not contain any TCL code, but it does use TCL to help with the
build process and to run tests.  You may need to install TCL development
libraries in order to successfully complete some makefile targets.
It is helpful, but is not required, to install the SQLite TCL extension
(the "tclextension-install" target) prior to running tests.</p>

<p>Build using Makefile.msc.  Example:</p>

<pre><code>    nmake /f Makefile.msc sqlite3.exe
    nmake /f Makefile.msc sqlite3.c
    nmake /f Makefile.msc sqldiff.exe
    # Makefile targets below this point require TCL development libraries
    nmake /f Makefile.msc tclextension-install
    nmake /f Makefile.msc devtest
    nmake /f Makefile.msc releasetest
    nmake /f Makefile.msc sqlite3_analyzer.exe
</code></pre>

<p>There are many other makefile targets.  See comments in Makefile.msc for
details.</p>

<p>As with the unix Makefile, the OPTIONS=... argument can be passed on the nmake
command-line to enable new compile-time options.  For example:</p>

<pre><code>    nmake /f Makefile.msc OPTIONS=-DSQLITE_OMIT_DEPRECATED sqlite3.exe
</code></pre>

<h2>Source Tree Map</h2>
<ul>
<li><p><strong>src/</strong> - This directory contains the primary source code for the
 SQLite core.  For historical reasons, C-code used for testing is
 also found here.  Source files intended for testing begin with "<code>test</code>".
 The <code>tclsqlite3.c</code> and <code>tclsqlite3.h</code> files are the TCL interface
 for SQLite and are also not part of the core.</p></li>
<li><p><strong>test/</strong> - This directory and its subdirectories contains code used
 for testing.  Files that end in "<code>.test</code>" are TCL scripts that run
 tests using an augmented TCL interpreter named "testfixture".  Use
 a command like "<code>make testfixture</code>" (unix) or 
 "<code>nmake /f Makefile.msc testfixture.exe</code>" (windows) to build that
 augmented TCL interpreter, then run individual tests using commands like
 "<code>testfixture test/main.test</code>".  This test/ subdirectory also contains
 additional C code modules and scripts for other kinds of testing.</p></li>
<li><p><strong>tool/</strong> - This directory contains programs and scripts used to
 build some of the machine-generated code that goes into the SQLite
 core, as well as to build and run tests and perform diagnostics.
 The source code to <a href="./doc/lemon.html">the Lemon parser generator</a> is
 found here.  There are also TCL scripts used to build and/or transform
 source code files.  For example, the tool/mksqlite3h.tcl script reads
 the src/sqlite.h.in file and uses it as a template to construct
 the deliverable "sqlite3.h" file that defines the SQLite interface.</p></li>
<li><p><strong>ext/</strong> - Various extensions to SQLite are found under this
 directory.  For example, the FTS5 subsystem is in "ext/fts5/".
 Some of these extensions (ex: FTS3/4, FTS5, RTREE) might get built
 into the SQLite amalgamation, but not all of them.  The
 "ext/misc/" subdirectory contains an assortment of one-file extensions,
 many of which are omitted from the SQLite core, but which are included
 in the <a href="https://sqlite.org/cli.html">SQLite CLI</a>.</p></li>
<li><p><strong>doc/</strong> - Some documentation files about SQLite internals are found
 here.  Note, however, that the primary documentation designed for
 application developers and users of SQLite is in a completely separate
 repository.  Note also that the primary API documentation is derived
 from specially constructed comments in the src/sqlite.h.in file.</p></li>
</ul>

<h3>Generated Source Code Files</h3>
<p>Several of the C-language source files used by SQLite are generated from
other sources rather than being typed in manually by a programmer.  This
section will summarize those automatically-generated files.  To create all
of the automatically-generated files, simply run "make target&#95;source".
The "target&#95;source" make target will create a subdirectory "tsrc/" and
fill it with all the source files needed to build SQLite, both
manually-edited files and automatically-generated files.</p>

<p>The SQLite interface is defined by the <strong>sqlite3.h</strong> header file, which is
generated from src/sqlite.h.in, ./manifest.uuid, and ./VERSION.  The
<a href="https://www.tcl.tk">Tcl script</a> at tool/mksqlite3h.tcl does the conversion.
The manifest.uuid file contains the SHA3 hash of the particular check-in
and is used to generate the SQLITE_SOURCE_ID macro.  The VERSION file
contains the current SQLite version number.  The sqlite3.h header is really
just a copy of src/sqlite.h.in with the source-id and version number inserted
at just the right spots. Note that comment text in the sqlite3.h file is
used to generate much of the SQLite API documentation.  The Tcl scripts
used to generate that documentation are in a separate source repository.</p>

<p>The SQL language parser is <strong>parse.c</strong> which is generated from a grammar in
the src/parse.y file.  The conversion of "parse.y" into "parse.c" is done
by the <a href="./doc/lemon.html">lemon</a> LALR(1) parser generator.  The source code
for lemon is at tool/lemon.c.  Lemon uses the tool/lempar.c file as a
template for generating its parser.
Lemon also generates the <strong>parse.h</strong> header file, at the same time it
generates parse.c.</p>

<p>The <strong>opcodes.h</strong> header file contains macros that define the numbers
corresponding to opcodes in the "VDBE" virtual machine.  The opcodes.h
file is generated by scanning the src/vdbe.c source file.  The
Tcl script at ./mkopcodeh.tcl does this scan and generates opcodes.h.
A second Tcl script, ./mkopcodec.tcl, then scans opcodes.h to generate
the <strong>opcodes.c</strong> source file, which contains a reverse mapping from
opcode-number to opcode-name that is used for EXPLAIN output.</p>

<p>The <strong>keywordhash.h</strong> header file contains the definition of a hash table
that maps SQL language keywords (ex: "CREATE", "SELECT", "INDEX", etc.) into
the numeric codes used by the parse.c parser.  The keywordhash.h file is
generated by a C-language program at tool mkkeywordhash.c.</p>

<p>The <strong>pragma.h</strong> header file contains various definitions used to parse
and implement the PRAGMA statements.  The header is generated by a
script <strong>tool/mkpragmatab.tcl</strong>. If you want to add a new PRAGMA, edit
the <strong>tool/mkpragmatab.tcl</strong> file to insert the information needed by the
parser for your new PRAGMA, then run the script to regenerate the
<strong>pragma.h</strong> header file.</p>

<h3>The Amalgamation</h3>
<p>All of the individual C source code and header files (both manually-edited
and automatically-generated) can be combined into a single big source file
<strong>sqlite3.c</strong> called "the amalgamation".  The amalgamation is the recommended
way of using SQLite in a larger application.  Combining all individual
source code files into a single big source code file allows the C compiler
to perform more cross-procedure analysis and generate better code.  SQLite
runs about 5% faster when compiled from the amalgamation versus when compiled
from individual source files.</p>

<p>The amalgamation is generated from the tool/mksqlite3c.tcl Tcl script.
First, all of the individual source files must be gathered into the tsrc/
subdirectory (using the equivalent of "make target_source") then the
tool/mksqlite3c.tcl script is run to copy them all together in just the
right order while resolving internal "#include" references.</p>

<p>The amalgamation source file is more than 200K lines long.  Some symbolic
debuggers (most notably MSVC) are unable to deal with files longer than 64K
lines.  To work around this, a separate Tcl script, tool/split-sqlite3c.tcl,
can be run on the amalgamation to break it up into a single small C file
called <strong>sqlite3-all.c</strong> that does #include on about seven other files
named <strong>sqlite3-1.c</strong>, <strong>sqlite3-2.c</strong>, ..., <strong>sqlite3-7.c</strong>.  In this way,
all of the source code is contained within a single translation unit so
that the compiler can do extra cross-procedure optimization, but no
individual source file exceeds 32K lines in length.</p>

<h2>How It All Fits Together</h2>
<p>SQLite is modular in design.
See the <a href="https://www.sqlite.org/arch.html">architectural description</a>
for details. Other documents that are useful in
helping to understand how SQLite works include the
<a href="https://www.sqlite.org/fileformat2.html">file format</a> description,
the <a href="https://www.sqlite.org/opcode.html">virtual machine</a> that runs
prepared statements, the description of
<a href="https://www.sqlite.org/atomiccommit.html">how transactions work</a>, and
the <a href="https://www.sqlite.org/optoverview.html">overview of the query planner</a>.</p>

<p>Decades of effort have gone into optimizing SQLite, both
for small size and high performance.  And optimizations tend to result in
complex code.  So there is a lot of complexity in the current SQLite
implementation.  It will not be the easiest library in the world to hack.</p>

<h3>Key source code files</h3>
<ul>
<li><p><strong>sqlite.h.in</strong> - This file defines the public interface to the SQLite
 library.  Readers will need to be familiar with this interface before
 trying to understand how the library works internally.  This file is
 really a template that is transformed into the "sqlite3.h" deliverable
 using a script invoked by the makefile.</p></li>
<li><p><strong>sqliteInt.h</strong> - this header file defines many of the data objects
 used internally by SQLite.  In addition to "sqliteInt.h", some
 subsystems inside of sQLite have their own header files.  These internal
 interfaces are not for use by applications.  They can and do change
 from one release of SQLite to the next.</p></li>
<li><p><strong>parse.y</strong> - This file describes the LALR(1) grammar that SQLite uses
 to parse SQL statements, and the actions that are taken at each step
 in the parsing process.  The file is processed by the
 <a href="./doc/lemon.html">Lemon Parser Generator</a> to produce the actual C code
 used for parsing.</p></li>
<li><p><strong>vdbe.c</strong> - This file implements the virtual machine that runs
 prepared statements.  There are various helper files whose names
 begin with "vdbe".  The VDBE has access to the vdbeInt.h header file
 which defines internal data objects.  The rest of SQLite interacts
 with the VDBE through an interface defined by vdbe.h.</p></li>
<li><p><strong>where.c</strong> - This file (together with its helper files named
 by "where*.c") analyzes the WHERE clause and generates
 virtual machine code to run queries efficiently.  This file is
 sometimes called the "query optimizer".  It has its own private
 header file, whereInt.h, that defines data objects used internally.</p></li>
<li><p><strong>btree.c</strong> - This file contains the implementation of the B-Tree
 storage engine used by SQLite.  The interface to the rest of the system
 is defined by "btree.h".  The "btreeInt.h" header defines objects
 used internally by btree.c and not published to the rest of the system.</p></li>
<li><p><strong>pager.c</strong> - This file contains the "pager" implementation, the
 module that implements transactions.  The "pager.h" header file
 defines the interface between pager.c and the rest of the system.</p></li>
<li><p><strong>os_unix.c</strong> and <strong>os_win.c</strong> - These two files implement the interface
 between SQLite and the underlying operating system using the run-time
 pluggable VFS interface.</p></li>
<li><p><strong>shell.c.in</strong> - This file is not part of the core SQLite library.  This
 is the file that, when linked against sqlite3.a, generates the
 "sqlite3.exe" command-line shell.  The "shell.c.in" file is transformed
 into "shell.c" as part of the build process.</p></li>
<li><p><strong>tclsqlite.c</strong> - This file implements the Tcl bindings for SQLite.  It
 is not part of the core SQLite library.  But as most of the tests in this
 repository are written in Tcl, the Tcl language bindings are important.</p></li>
<li><p><strong>test*.c</strong> - Files in the src/ folder that begin with "test" go into
 building the "testfixture.exe" program.  The testfixture.exe program is
 an enhanced Tcl shell.  The testfixture.exe program runs scripts in the
 test/ folder to validate the core SQLite code.  The testfixture program
 (and some other test programs too) is built and run when you type
 "make test".</p></li>
<li><p><strong>VERSION</strong>, <strong>manifest</strong>, and <strong>manifest.uuid</strong> - These files define
 the current SQLite version number.  The "VERSION" file is human generated,
 but the "manifest" and "manifest.uuid" files are automatically generated
 by the <a href="https://fossil-scm.org/">Fossil version control system</a>.</p></li>
</ul>

<p>There are many other source files.  Each has a succinct header comment that
describes its purpose and role within the larger system.</p>

<p><a name="vauth"></a></p>

<h2>Verifying Code Authenticity</h2>
<p>The <code>manifest</code> file at the root directory of the source tree
contains either a SHA3-256 hash or a SHA1 hash
for every source file in the repository.
The name of the version of the entire source tree is just the
SHA3-256 hash of the <code>manifest</code> file itself, possibly with the
last line of that file omitted if the last line begins with
"<code># Remove this line</code>".
The <code>manifest.uuid</code> file should contain the SHA3-256 hash of the
<code>manifest</code> file. If all of the above hash comparisons are correct, then
you can be confident that your source tree is authentic and unadulterated.
Details on the format for the <code>manifest</code> files are available
<a href="https://fossil-scm.org/home/doc/trunk/www/fileformat.wiki#manifest">on the Fossil website</a>.</p>

<p>The process of checking source code authenticity is automated by the 
makefile:</p>

<blockquote>
<p>  make verify-source</p>
</blockquote>

<p>Or on windows:</p>

<blockquote>
<p>  nmake /f Makefile.msc verify-source</p>
</blockquote>

<p>Using the makefile to verify source integrity is good for detecting
accidental changes to the source tree, but malicious changes could be
hidden by also modifying the makefiles.</p>

<h2>Contacts</h2>
<p>The main SQLite website is <a href="https://sqlite.org/">https://sqlite.org/</a>
with geographically distributed backups at
<a href="https://www2.sqlite.org">https://www2.sqlite.org/</a> and
<a href="https://www3.sqlite.org">https://www3.sqlite.org/</a>.</p>

<p>Contact the SQLite developers through the
<a href="https://sqlite.org/forum/">SQLite Forum</a>.  In an emergency, you
can send private email to the lead developer at drh at sqlite dot org.</p>

</div>
<script nonce='5c468789942c60a9e666383658ec74ab1612bc120dd0b847'>/* builtin.c:621 */
(function(){
if(window.NodeList && !NodeList.prototype.forEach){NodeList.prototype.forEach = Array.prototype.forEach;}
if(!window.fossil) window.fossil={};
window.fossil.version = "2.26 [680acb2831] 2025-01-06 20:52:56 UTC";
window.fossil.rootPath = "/src"+'/';
window.fossil.config = {projectName: "SQLite",
shortProjectName: "",
projectCode: "2ab58778c2967968b94284e989e43dc11791f548",
/* Length of UUID hashes for display purposes. */hashDigits: 8, hashDigitsUrl: 16,
diffContextLines: 5,
editStateMarkers: {/*Symbolic markers to denote certain edit states.*/isNew:'[+]', isModified:'[*]', isDeleted:'[-]'},
confirmerButtonTicks: 3 /*default fossil.confirmer tick count.*/,
skin:{isDark: false/*true if the current skin has the 'white-foreground' detail*/}
};
window.fossil.user = {name: "guest",isAdmin: false};
if(fossil.config.skin.isDark) document.body.classList.add('fossil-dark-style');
window.fossil.page = {name:"doc/trunk/README.md"};
})();
</script>
<script nonce='5c468789942c60a9e666383658ec74ab1612bc120dd0b847'>/* doc.c:434 */
window.addEventListener('load', ()=>window.fossil.pikchr.addSrcView(), false);
</script>
</div>
<footer>
This page was generated in about
0.013s by
Fossil 2.26 [680acb2831] 2025-01-06 20:52:56
</footer>
<script nonce="5c468789942c60a9e666383658ec74ab1612bc120dd0b847">/* style.c:898 */
function debugMsg(msg){
var n = document.getElementById("debugMsg");
if(n){n.textContent=msg;}
}
</script>
<script nonce='5c468789942c60a9e666383658ec74ab1612bc120dd0b847'>
/* hbmenu.js *************************************************************/
(function() {
var hbButton = document.getElementById("hbbtn");
if (!hbButton) return;
if (!document.addEventListener) return;
var panel = document.getElementById("hbdrop");
if (!panel) return;
if (!panel.style) return;
var panelBorder = panel.style.border;
var panelInitialized = false;
var panelResetBorderTimerID = 0;
var animate = panel.style.transition !== null && (typeof(panel.style.transition) == "string");
var animMS = panel.getAttribute("data-anim-ms");
if (animMS) {
animMS = parseInt(animMS);
if (isNaN(animMS) || animMS == 0)
animate = false;
else if (animMS < 0)
animMS = 400;
}
else
animMS = 400;
var panelHeight;
function calculatePanelHeight() {
panel.style.maxHeight = '';
var es   = window.getComputedStyle(panel),
edis = es.display,
epos = es.position,
evis = es.visibility;
panel.style.visibility = 'hidden';
panel.style.position   = 'absolute';
panel.style.display    = 'block';
panelHeight = panel.offsetHeight + 'px';
panel.style.display    = edis;
panel.style.position   = epos;
panel.style.visibility = evis;
}
function showPanel() {
if (panelResetBorderTimerID) {
clearTimeout(panelResetBorderTimerID);
panelResetBorderTimerID = 0;
}
if (animate) {
if (!panelInitialized) {
panelInitialized = true;
calculatePanelHeight();
panel.style.transition = 'max-height ' + animMS +
'ms ease-in-out';
panel.style.overflowY  = 'hidden';
panel.style.maxHeight  = '0';
}
setTimeout(function() {
panel.style.maxHeight = panelHeight;
panel.style.border    = panelBorder;
}, 40);
}
panel.style.display = 'block';
document.addEventListener('keydown',panelKeydown,true);
document.addEventListener('click',panelClick,false);
}
var panelKeydown = function(event) {
var key = event.which || event.keyCode;
if (key == 27) {
event.stopPropagation();
panelToggle(true);
}
};
var panelClick = function(event) {
if (!panel.contains(event.target)) {
panelToggle(true);
}
};
function panelShowing() {
if (animate) {
return panel.style.maxHeight == panelHeight;
}
else {
return panel.style.display == 'block';
}
}
function hasChildren(element) {
var childElement = element.firstChild;
while (childElement) {
if (childElement.nodeType == 1)
return true;
childElement = childElement.nextSibling;
}
return false;
}
window.addEventListener('resize',function(event) {
panelInitialized = false;
},false);
hbButton.addEventListener('click',function(event) {
event.stopPropagation();
event.preventDefault();
panelToggle(false);
},false);
function panelToggle(suppressAnimation) {
if (panelShowing()) {
document.removeEventListener('keydown',panelKeydown,true);
document.removeEventListener('click',panelClick,false);
if (animate) {
if (suppressAnimation) {
var transition = panel.style.transition;
panel.style.transition = '';
panel.style.maxHeight = '0';
panel.style.border = 'none';
setTimeout(function() {
panel.style.transition = transition;
}, 40);
}
else {
panel.style.maxHeight = '0';
panelResetBorderTimerID = setTimeout(function() {
panel.style.border = 'none';
panelResetBorderTimerID = 0;
}, animMS);
}
}
else {
panel.style.display = 'none';
}
}
else {
if (!hasChildren(panel)) {
var xhr = new XMLHttpRequest();
xhr.onload = function() {
var doc = xhr.responseXML;
if (doc) {
var sm = doc.querySelector("ul#sitemap");
if (sm && xhr.status == 200) {
panel.innerHTML = sm.outerHTML;
showPanel();
}
}
}
var url = hbButton.href + (hbButton.href.includes("?")?"&popup":"?popup")
xhr.open("GET", url);
xhr.responseType = "document";
xhr.send();
}
else {
showPanel();
}
}
}
})();
/* fossil.bootstrap.js *************************************************************/
"use strict";
(function () {
if(typeof window.CustomEvent === "function") return false;
window.CustomEvent = function(event, params) {
if(!params) params = {bubbles: false, cancelable: false, detail: null};
const evt = document.createEvent('CustomEvent');
evt.initCustomEvent( event, !!params.bubbles, !!params.cancelable, params.detail );
return evt;
};
})();
(function(global){
const F = global.fossil;
const timestring = function f(){
if(!f.rx1){
f.rx1 = /\.\d+Z$/;
}
const d = new Date();
return d.toISOString().replace(f.rx1,'').split('T').join(' ');
};
const localTimeString = function ff(d){
if(!ff.pad){
ff.pad = (x)=>(''+x).length>1 ? x : '0'+x;
}
d || (d = new Date());
return [
d.getFullYear(),'-',ff.pad(d.getMonth()+1),
'-',ff.pad(d.getDate()),
' ',ff.pad(d.getHours()),':',ff.pad(d.getMinutes()),
':',ff.pad(d.getSeconds())
].join('');
};
F.message = function f(msg){
const args = Array.prototype.slice.call(arguments,0);
const tgt = f.targetElement;
if(args.length) args.unshift(
localTimeString()+':'
);
if(tgt){
tgt.classList.remove('error');
tgt.innerText = args.join(' ');
}
else{
if(args.length){
args.unshift('Fossil status:');
console.debug.apply(console,args);
}
}
return this;
};
F.message.targetElement =
document.querySelector('#fossil-status-bar');
if(F.message.targetElement){
F.message.targetElement.addEventListener(
'dblclick', ()=>F.message(), false
);
}
F.error = function f(msg){
const args = Array.prototype.slice.call(arguments,0);
const tgt = F.message.targetElement;
args.unshift(timestring(),'UTC:');
if(tgt){
tgt.classList.add('error');
tgt.innerText = args.join(' ');
}
else{
args.unshift('Fossil error:');
console.error.apply(console,args);
}
return this;
};
F.encodeUrlArgs = function(obj,tgtArray,fakeEncode){
if(!obj) return '';
const a = (tgtArray instanceof Array) ? tgtArray : [],
enc = fakeEncode ? (x)=>x : encodeURIComponent;
let k, i = 0;
for( k in obj ){
if(i++) a.push('&');
a.push(enc(k),'=',enc(obj[k]));
}
return a===tgtArray ? a : a.join('');
};
F.repoUrl = function(path,urlParams){
if(!urlParams) return this.rootPath+path;
const url=[this.rootPath,path];
url.push('?');
if('string'===typeof urlParams) url.push(urlParams);
else if(urlParams && 'object'===typeof urlParams){
this.encodeUrlArgs(urlParams, url);
}
return url.join('');
};
F.isObject = function(v){
return v &&
(v instanceof Object) &&
('[object Object]' === Object.prototype.toString.apply(v) );
};
F.mergeLastWins = function(){
var k, o, i;
const n = arguments.length, rc={};
for(i = 0; i < n; ++i){
if(!F.isObject(o = arguments[i])) continue;
for( k in o ){
if(o.hasOwnProperty(k)) rc[k] = o[k];
}
}
return rc;
};
F.hashDigits = function(hash,forUrl){
const n = ('number'===typeof forUrl)
? forUrl : F.config[forUrl ? 'hashDigitsUrl' : 'hashDigits'];
return ('string'==typeof hash ? hash.substr(
0, n
) : hash);
};
F.onPageLoad = function(callback){
window.addEventListener('load', callback, false);
return this;
};
F.onDOMContentLoaded = function(callback){
window.addEventListener('DOMContentLoaded', callback, false);
return this;
};
F.shortenFilename = function(name){
const a = name.split('/');
if(a.length<=2) return name;
while(a.length>2) a.shift();
return '.../'+a.join('/');
};
F.page.addEventListener = function f(eventName, callback){
if(!f.proxy){
f.proxy = document.createElement('span');
}
f.proxy.addEventListener(eventName, callback, false);
return this;
};
F.page.dispatchEvent = function(eventName, eventDetail){
if(this.addEventListener.proxy){
try{
this.addEventListener.proxy.dispatchEvent(
new CustomEvent(eventName,{detail: eventDetail})
);
}catch(e){
console.error(eventName,"event listener threw:",e);
}
}
return this;
};
F.page.setPageTitle = function(title){
const t = document.querySelector('title');
if(t) t.innerText = title;
return this;
};
F.debounce = function f(func, waitMs, immediate) {
var timeoutId;
if(!waitMs) waitMs = f.$defaultDelay;
return function() {
const context = this, args = Array.prototype.slice.call(arguments);
const later = function() {
timeoutId = undefined;
if(!immediate) func.apply(context, args);
};
const callNow = immediate && !timeoutId;
clearTimeout(timeoutId);
timeoutId = setTimeout(later, waitMs);
if(callNow) func.apply(context, args);
};
};
F.debounce.$defaultDelay = 500;
})(window);
/* fossil.dom.js *************************************************************/
"use strict";
(function(F){
const argsToArray = (a)=>Array.prototype.slice.call(a,0);
const isArray = (v)=>v instanceof Array;
const dom = {
create: function(elemType){
return document.createElement(elemType);
},
createElemFactory: function(eType){
return function(){
return document.createElement(eType);
};
},
remove: function(e){
if(e.forEach){
e.forEach(
(x)=>x.parentNode.removeChild(x)
);
}else{
e.parentNode.removeChild(e);
}
return e;
},
clearElement: function f(e){
if(!f.each){
f.each = function(e){
if(e.forEach){
e.forEach((x)=>f(x));
return e;
}
while(e.firstChild) e.removeChild(e.firstChild);
};
}
argsToArray(arguments).forEach(f.each);
return arguments[0];
},
};
dom.splitClassList = function f(str){
if(!f.rx){
f.rx = /(\s+|\s*,\s*)/;
}
return str ? str.split(f.rx) : [str];
};
dom.div = dom.createElemFactory('div');
dom.p = dom.createElemFactory('p');
dom.code = dom.createElemFactory('code');
dom.pre = dom.createElemFactory('pre');
dom.header = dom.createElemFactory('header');
dom.footer = dom.createElemFactory('footer');
dom.section = dom.createElemFactory('section');
dom.span = dom.createElemFactory('span');
dom.strong = dom.createElemFactory('strong');
dom.em = dom.createElemFactory('em');
dom.ins = dom.createElemFactory('ins');
dom.del = dom.createElemFactory('del');
dom.label = function(forElem, text){
const rc = document.createElement('label');
if(forElem){
if(forElem instanceof HTMLElement){
forElem = this.attr(forElem, 'id');
}
if(forElem){
dom.attr(rc, 'for', forElem);
}
}
if(text) this.append(rc, text);
return rc;
};
dom.img = function(src){
const e = this.create('img');
if(src) e.setAttribute('src',src);
return e;
};
dom.a = function(href,label){
const e = this.create('a');
if(href) e.setAttribute('href',href);
if(label) e.appendChild(dom.text(true===label ? href : label));
return e;
};
dom.hr = dom.createElemFactory('hr');
dom.br = dom.createElemFactory('br');
dom.text = function(){
return document.createTextNode(argsToArray(arguments).join(''));
};
dom.button = function(label,callback){
const b = this.create('button');
if(label) b.appendChild(this.text(label));
if('function' === typeof callback){
b.addEventListener('click', callback, false);
}
return b;
};
dom.textarea = function(){
const rc = this.create('textarea');
let rows, cols, readonly;
if(1===arguments.length){
if('boolean'===typeof arguments[0]){
readonly = !!arguments[0];
}else{
rows = arguments[0];
}
}else if(arguments.length){
rows = arguments[0];
cols = arguments[1];
readonly = arguments[2];
}
if(rows) rc.setAttribute('rows',rows);
if(cols) rc.setAttribute('cols', cols);
if(readonly) rc.setAttribute('readonly', true);
return rc;
};
dom.select = dom.createElemFactory('select');
dom.option = function(value,label){
const a = arguments;
var sel;
if(1==a.length){
if(a[0] instanceof HTMLElement){
sel = a[0];
}else{
value = a[0];
}
}else if(2==a.length){
if(a[0] instanceof HTMLElement){
sel = a[0];
value = a[1];
}else{
value = a[0];
label = a[1];
}
}
else if(3===a.length){
sel = a[0];
value = a[1];
label = a[2];
}
const o = this.create('option');
if(undefined !== value){
o.value = value;
this.append(o, this.text(label || value));
}else if(undefined !== label){
this.append(o, label);
}
if(sel) this.append(sel, o);
return o;
};
dom.h = function(level){
return this.create('h'+level);
};
dom.ul = dom.createElemFactory('ul');
dom.li = function(parent){
const li = this.create('li');
if(parent) parent.appendChild(li);
return li;
};
dom.createElemFactoryWithOptionalParent = function(childType){
return function(parent){
const e = this.create(childType);
if(parent) parent.appendChild(e);
return e;
};
};
dom.table = dom.createElemFactory('table');
dom.thead = dom.createElemFactoryWithOptionalParent('thead');
dom.tbody = dom.createElemFactoryWithOptionalParent('tbody');
dom.tfoot = dom.createElemFactoryWithOptionalParent('tfoot');
dom.tr = dom.createElemFactoryWithOptionalParent('tr');
dom.td = dom.createElemFactoryWithOptionalParent('td');
dom.th = dom.createElemFactoryWithOptionalParent('th');
dom.fieldset = function(legendText){
const fs = this.create('fieldset');
if(legendText){
this.append(
fs,
(legendText instanceof HTMLElement)
? legendText
: this.append(this.legend(legendText))
);
}
return fs;
};
dom.legend = function(legendText){
const rc = this.create('legend');
if(legendText) this.append(rc, legendText);
return rc;
};
dom.append = function f(parent){
const a = argsToArray(arguments);
a.shift();
for(let i in a) {
var e = a[i];
if(isArray(e) || e.forEach){
e.forEach((x)=>f.call(this, parent,x));
continue;
}
if('string'===typeof e
|| 'number'===typeof e
|| 'boolean'===typeof e
|| e instanceof Error) e = this.text(e);
parent.appendChild(e);
}
return parent;
};
dom.input = function(type){
return this.attr(this.create('input'), 'type', type);
};
dom.checkbox = function(value, checked){
const rc = this.input('checkbox');
if(1===arguments.length && 'boolean'===typeof value){
checked = !!value;
value = undefined;
}
if(undefined !== value) rc.value = value;
if(!!checked) rc.checked = true;
return rc;
};
dom.radio = function(){
const rc = this.input('radio');
let name, value, checked;
if(1===arguments.length && 'boolean'===typeof name){
checked = arguments[0];
name = value = undefined;
}else if(2===arguments.length){
name = arguments[0];
if('boolean'===typeof arguments[1]){
checked = arguments[1];
}else{
value = arguments[1];
checked = undefined;
}
}else if(arguments.length){
name = arguments[0];
value = arguments[1];
checked = arguments[2];
}
if(name) this.attr(rc, 'name', name);
if(undefined!==value) rc.value = value;
if(!!checked) rc.checked = true;
return rc;
};
const domAddRemoveClass = function f(action,e){
if(!f.rxSPlus){
f.rxSPlus = /\s+/;
f.applyAction = function(e,a,v){
if(!e || !v
) return;
else if(e.forEach){
e.forEach((E)=>E.classList[a](v));
}else{
e.classList[a](v);
}
};
}
var i = 2, n = arguments.length;
for( ; i < n; ++i ){
let c = arguments[i];
if(!c) continue;
else if(isArray(c) ||
('string'===typeof c
&& c.indexOf(' ')>=0
&& (c = c.split(f.rxSPlus)))
|| c.forEach
){
c.forEach((k)=>k ? f.applyAction(e, action, k) : false);
}else if(c){
f.applyAction(e, action, c);
}
}
return e;
};
dom.addClass = function(e,c){
const a = argsToArray(arguments);
a.unshift('add');
return domAddRemoveClass.apply(this, a);
};
dom.removeClass = function(e,c){
const a = argsToArray(arguments);
a.unshift('remove');
return domAddRemoveClass.apply(this, a);
};
dom.toggleClass = function f(e,c){
if(e.forEach){
e.forEach((x)=>x.classList.toggle(c));
}else{
e.classList.toggle(c);
}
return e;
};
dom.hasClass = function(e,c){
return (e && e.classList) ? e.classList.contains(c) : false;
};
dom.moveTo = function(dest,e){
const n = arguments.length;
var i = 1;
const self = this;
for( ; i < n; ++i ){
e = arguments[i];
this.append(dest, e);
}
return dest;
};
dom.moveChildrenTo = function f(dest,e){
if(!f.mv){
f.mv = function(d,v){
if(d instanceof Array){
d.push(v);
if(v.parentNode) v.parentNode.removeChild(v);
}
else d.appendChild(v);
};
}
const n = arguments.length;
var i = 1;
for( ; i < n; ++i ){
e = arguments[i];
if(!e){
console.warn("Achtung: dom.moveChildrenTo() passed a falsy value at argument",i,"of",
arguments,arguments[i]);
continue;
}
if(e.forEach){
e.forEach((x)=>f.mv(dest, x));
}else{
while(e.firstChild){
f.mv(dest, e.firstChild);
}
}
}
return dest;
};
dom.replaceNode = function f(old,nu){
var i = 1, n = arguments.length;
++f.counter;
try {
for( ; i < n; ++i ){
const e = arguments[i];
if(e.forEach){
e.forEach((x)=>f.call(this,old,e));
continue;
}
old.parentNode.insertBefore(e, old);
}
}
finally{
--f.counter;
}
if(!f.counter){
old.parentNode.removeChild(old);
}
};
dom.replaceNode.counter = 0;
dom.attr = function f(e){
if(2===arguments.length) return e.getAttribute(arguments[1]);
const a = argsToArray(arguments);
if(e.forEach){
e.forEach(function(x){
a[0] = x;
f.apply(f,a);
});
return e;
}
a.shift();
while(a.length){
const key = a.shift(), val = a.shift();
if(null===val || undefined===val){
e.removeAttribute(key);
}else{
e.setAttribute(key,val);
}
}
return e;
};
const enableDisable = function f(enable){
var i = 1, n = arguments.length;
for( ; i < n; ++i ){
let e = arguments[i];
if(e.forEach){
e.forEach((x)=>f(enable,x));
}else{
e.disabled = !enable;
}
}
return arguments[1];
};
dom.enable = function(e){
const args = argsToArray(arguments);
args.unshift(true);
return enableDisable.apply(this,args);
};
dom.disable = function(e){
const args = argsToArray(arguments);
args.unshift(false);
return enableDisable.apply(this,args);
};
dom.selectOne = function(x,origin){
var src = origin || document,
e = src.querySelector(x);
if(!e){
e = new Error("Cannot find DOM element: "+x);
console.error(e, src);
throw e;
}
return e;
};
dom.flashOnce = function f(e,howLongMs,afterFlashCallback){
if(e.dataset.isBlinking){
return;
}
if(2===arguments.length && 'function' ===typeof howLongMs){
afterFlashCallback = howLongMs;
howLongMs = f.defaultTimeMs;
}
if(!howLongMs || 'number'!==typeof howLongMs){
howLongMs = f.defaultTimeMs;
}
e.dataset.isBlinking = true;
const transition = e.style.transition;
e.style.transition = "opacity "+howLongMs+"ms ease-in-out";
const opacity = e.style.opacity;
e.style.opacity = 0;
setTimeout(function(){
e.style.transition = transition;
e.style.opacity = opacity;
delete e.dataset.isBlinking;
if(afterFlashCallback) afterFlashCallback();
}, howLongMs);
return e;
};
dom.flashOnce.defaultTimeMs = 400;
dom.flashOnce.eventHandler = (event)=>dom.flashOnce(event.target)
dom.flashNTimes = function(e,n,howLongMs,afterFlashCallback){
const args = argsToArray(arguments);
args.splice(1,1);
if(arguments.length===3 && 'function'===typeof howLongMs){
afterFlashCallback = howLongMs;
howLongMs = args[1] = this.flashOnce.defaultTimeMs;
}else if(arguments.length<3){
args[1] = this.flashOnce.defaultTimeMs;
}
n = +n;
const self = this;
const cb = args[2] = function f(){
if(--n){
setTimeout(()=>self.flashOnce(e, howLongMs, f),
howLongMs+(howLongMs*0.1));
}else if(afterFlashCallback){
afterFlashCallback();
}
};
this.flashOnce.apply(this, args);
return this;
};
dom.addClassBriefly = function f(e, className, howLongMs, afterCallback){
if(arguments.length<4 && 'function'===typeof howLongMs){
afterCallback = howLongMs;
howLongMs = f.defaultTimeMs;
}else if(arguments.length<3 || !+howLongMs){
howLongMs = f.defaultTimeMs;
}
this.addClass(e, className);
setTimeout(function(){
dom.removeClass(e, className);
if(afterCallback) afterCallback();
}, howLongMs);
return this;
};
dom.addClassBriefly.defaultTimeMs = 1000;
dom.copyTextToClipboard = function(text){
if( window.clipboardData && window.clipboardData.setData ){
window.clipboardData.setData('Text',text);
return true;
}else{
const x = document.createElement("textarea");
x.style.position = 'fixed';
x.value = text;
document.body.appendChild(x);
x.select();
var rc;
try{
document.execCommand('copy');
rc = true;
}catch(err){
rc = false;
}finally{
document.body.removeChild(x);
}
return rc;
}
};
dom.copyStyle = function f(e, style){
if(e.forEach){
e.forEach((x)=>f(x, style));
return e;
}
if(style){
let k;
for(k in style){
if(style.hasOwnProperty(k)) e.style[k] = style[k];
}
}
return e;
};
dom.effectiveHeight = function f(e){
if(!e) return 0;
if(!f.measure){
f.measure = function callee(e, depth){
if(!e) return;
const m = e.getBoundingClientRect();
if(0===depth){
callee.top = m.top;
callee.bottom = m.bottom;
}else{
callee.top = m.top ? Math.min(callee.top, m.top) : callee.top;
callee.bottom = Math.max(callee.bottom, m.bottom);
}
Array.prototype.forEach.call(e.children,(e)=>callee(e,depth+1));
if(0===depth){
f.extra += callee.bottom - callee.top;
}
return f.extra;
};
}
f.extra = 0;
f.measure(e,0);
return f.extra;
};
dom.parseHtml = function(){
let childs, string, tgt;
if(1===arguments.length){
string = arguments[0];
}else if(2==arguments.length){
tgt = arguments[0];
string  = arguments[1];
}
if(string){
const newNode = new DOMParser().parseFromString(string, 'text/html');
childs = newNode.documentElement.querySelector('body');
childs = childs ? Array.prototype.slice.call(childs.childNodes, 0) : [];
}else{
childs = [];
}
return tgt ? this.moveTo(tgt, childs) : childs;
};
F.connectPagePreviewers = function f(selector,methodNamespace){
if('string'===typeof selector){
selector = document.querySelectorAll(selector);
}else if(!selector.forEach){
selector = [selector];
}
if(!methodNamespace){
methodNamespace = F.page;
}
selector.forEach(function(e){
e.addEventListener(
'click', function(r){
const eTo = '#'===e.dataset.fPreviewTo[0]
? document.querySelector(e.dataset.fPreviewTo)
: methodNamespace[e.dataset.fPreviewTo],
eFrom = '#'===e.dataset.fPreviewFrom[0]
? document.querySelector(e.dataset.fPreviewFrom)
: methodNamespace[e.dataset.fPreviewFrom],
asText = +(e.dataset.fPreviewAsText || 0);
eTo.textContent = "Fetching preview...";
methodNamespace[e.dataset.fPreviewVia](
(eFrom instanceof Function ? eFrom.call(methodNamespace) : eFrom.value),
function(r){
if(eTo instanceof Function) eTo.call(methodNamespace, r||'');
else if(!r){
dom.clearElement(eTo);
}else if(asText){
eTo.textContent = r;
}else{
dom.parseHtml(dom.clearElement(eTo), r);
}
}
);
}, false
);
});
return this;
};
return F.dom = dom;
})(window.fossil);
/* fossil.pikchr.js *************************************************************/
(function(F){
"use strict";
const D = F.dom, P = F.pikchr = {};
P.addSrcView = function f(svg){
if(!f.hasOwnProperty('parentClick')){
f.parentClick = function(ev){
if(ev.altKey || ev.metaKey || ev.ctrlKey
|| this.classList.contains('toggle')){
this.classList.toggle('source');
ev.stopPropagation();
ev.preventDefault();
}
};
f.clickPikchrShow = function(ev){
const pId = this.dataset['pikchrid'];
if(!pId) return;
const ePikchr = this.parentNode.parentNode.querySelector('#'+pId);
if(!ePikchr) return;
ev.stopPropagation();
window.sessionStorage.setItem('pikchr-xfer', ePikchr.innerText);
};
};
if(!svg) svg = 'svg.pikchr';
if('string' === typeof svg){
document.querySelectorAll(svg).forEach((e)=>f.call(this, e));
return this;
}else if(svg.forEach){
svg.forEach((e)=>f.call(this, e));
return this;
}
if(svg.dataset.pikchrProcessed){
return this;
}
svg.dataset.pikchrProcessed = 1;
const parent = svg.parentNode.parentNode;
const srcView = parent ? svg.parentNode.nextElementSibling : undefined;
if(srcView && srcView.classList.contains('pikchr-src')){
parent.addEventListener('click', f.parentClick, false);
const eSpan = window.sessionStorage
? srcView.querySelector('span')
: undefined;
if(eSpan){
const openLink = eSpan.querySelector('a');
if(openLink){
openLink.addEventListener('click', f.clickPikchrShow, false);
eSpan.classList.remove('hidden');
}
}
}
return this;
};
})(window.fossil);
</script>
</body>
</html>
