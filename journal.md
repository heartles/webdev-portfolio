# Journal

Jane Mellott

Deployed site: https://heartles.io

Github Repo: https://github.com/heartles/webdev-portfolio

---

I decided early on I wanted to have a navbar that turned into a
hamburger menu toggle at low resolutions. I had enough experience with
making a navbar that I could display all the links, but I wasn't exactly
sure how to implement the hamburger menu. I consulted the internet and
found a tutorial from w3schools.com that I was able to adapt into the
navbar. This actually taught me about media queries, before I saw them
in class.

I also had a bit of trouble finding the right size that looked good for
the social media icon links. Most of that, however, was due to some CSS
issues that had been hiding. I bumped up the font size and used that
to get the right icon size.

---

For the first bit of the portfolio due as part of assignment 1, I had
already written well over 75 lines of CSS by the time I had the navbar
working. The links weren't connected yet, but as I had no pages to
connect them to, I wasn't too worried about that. When looking over
the requirements before turning in my work, I realized that there were
a whole set of requirements that had somehow slipped my mind. While the
nav bar was good, I also needed a short bio, some images, and a footer
still. When trying to put it together, I ran into a few issues.

### Issue 1

It was hard for me to write a bio that didn't feel either useless or
cliche'd. Usually when writing bios, I can get away with just linking
some of my social media profiles, but I had already committed to putting
those in the navbar. I ended up coming up with a few bullet points and
listing them out. I would eventually like to add some more content to the
home page however.

### Issue 2

While putting the footer at the bottom of the page or screen was simple,
getting the footer to extend down to the bottom of the screen was more
difficult. After spending a good amount of time trying to make the footer
extend if needed, I realized that I was doing this the hard way, and it
would be much simpler (and probably more correct) to just make the
background color of the page as a whole the same as the footer's.
This meant that I had to reframe my internal thoughts on the site,
where the header and footer "wrapped" the main section, and sat on
top of it, so that the body instead was "laid on top" of the header
and footer.

### Issue 3

I had issues with formatting the main section of the website (between
the header and footer). I wanted it to have a maximum width as a set
pixel count on larger screens, but still maintain a minimum distance
from the edges of the screen. I also needed the background color to apply
to the entire space between the header and footer, ignoring the margins.
I had issues with getting padding to work how I wanted it to (often the
content at the top of the page would press up against the header), so
my eventual solution was to have a sequence of three nested div's. While
I do think this is a bit janky, and would be an eventual target for
refactor, I do have it working now, and want to avoid accidentally
breaking it when I have to turn the site in.

---

I was able to get a basic projects page going using some of the work I
had done for the assignment 1 portion of the portfolio, but I had
a problem with finding the right content. I had a few projects that I
was currently working on, but few of them were complete enough that
I felt comfortable sharing them. I eventually put my full stack final
project there as a "in progress" item, the last big project I worked on,
and a group project I maintain with my friends that I contribute to.
I also later put in a few of my (more complete) old projects under a separate
heading.

---

I knew that, as the number of pages grew, duplicating the header and footer
(and maintaining that when it changed) was going to cause me significant
headache. I decided to find some sort of templating engine that I could
use, but as I was doing research they all looked a bit too heavy duty
for my use case. In the end, I wrote a very small go program that
prepared my site for deployment by injecting html files into a base file
that contained the header and footer. This produced a folder that
had exclusively static files, which I then used for deployment on Firebase.
I also put the compiled program in the repository (for Linux x86_64) to
make CI/CD easier to use.

I attempted to setup CI/CD through Github, but only got halfway done.
Since deployment was a single line command from my pc, I was not in a
huge rush to fully setup my pipeline.

---

I realized at this point that I hadn't done anything with bootstrap, and
the requirements implied that I needed to use it within my site. I decided
to make the contact form with bootstrap (modifying a few things to make the
design fit my site a bit better). I used my assignment 2 form as a base, then
edited the colors to fit the theme of the site (checking accessibility with
Axe) as well as the font. I decided to leave the form somewhat distinct from
the rest of the site, as I'm probably going to remove it after this term ends.
There's also a "return to site" link to further drive home the separation
of this contact form from the main site.
