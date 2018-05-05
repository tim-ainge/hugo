// Copyright 2018 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file is autogenerated.

// Package embedded defines the internal templates that Hugo provides.
package embedded

var EmbeddedTemplates = [][2]string{
	{`README.md`, `

## Build Templates

If you add or modify any template in this folder, you also need to run **mage generate** to get the Go code in synch.
`},
	{`_default/robots.txt`, `User-agent: *`},
	{`_default/rss.xml`, `<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
  <channel>
    <title>{{ if eq  .Title  .Site.Title }}{{ .Site.Title }}{{ else }}{{ with .Title }}{{.}} on {{ end }}{{ .Site.Title }}{{ end }}</title>
    <link>{{ .Permalink }}</link>
    <description>Recent content {{ if ne  .Title  .Site.Title }}{{ with .Title }}in {{.}} {{ end }}{{ end }}on {{ .Site.Title }}</description>
    <generator>Hugo -- gohugo.io</generator>{{ with .Site.LanguageCode }}
    <language>{{.}}</language>{{end}}{{ with .Site.Author.email }}
    <managingEditor>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</managingEditor>{{end}}{{ with .Site.Author.email }}
    <webMaster>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</webMaster>{{end}}{{ with .Site.Copyright }}
    <copyright>{{.}}</copyright>{{end}}{{ if not .Date.IsZero }}
    <lastBuildDate>{{ .Date.Format "Mon, 02 Jan 2006 15:04:05 -0700" | safeHTML }}</lastBuildDate>{{ end }}
    {{ with .OutputFormats.Get "RSS" }}
	{{ printf "<atom:link href=%q rel=\"self\" type=%q />" .Permalink .MediaType | safeHTML }}
    {{ end }}
    {{ range .Data.Pages }}
    <item>
      <title>{{ .Title }}</title>
      <link>{{ .Permalink }}</link>
      <pubDate>{{ .Date.Format "Mon, 02 Jan 2006 15:04:05 -0700" | safeHTML }}</pubDate>
      {{ with .Site.Author.email }}<author>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</author>{{end}}
      <guid>{{ .Permalink }}</guid>
      <description>{{ .Summary | html }}</description>
    </item>
    {{ end }}
  </channel>
</rss>`},
	{`_default/sitemap.xml`, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
  xmlns:xhtml="http://www.w3.org/1999/xhtml">
  {{ range .Data.Pages }}
  <url>
    <loc>{{ .Permalink }}</loc>{{ if not .Lastmod.IsZero }}
    <lastmod>{{ safeHTML ( .Lastmod.Format "2006-01-02T15:04:05-07:00" ) }}</lastmod>{{ end }}{{ with .Sitemap.ChangeFreq }}
    <changefreq>{{ . }}</changefreq>{{ end }}{{ if ge .Sitemap.Priority 0.0 }}
    <priority>{{ .Sitemap.Priority }}</priority>{{ end }}{{ if .IsTranslated }}{{ range .Translations }}
    <xhtml:link
                rel="alternate"
                hreflang="{{ .Lang }}"
                href="{{ .Permalink }}"
                />{{ end }}
    <xhtml:link
                rel="alternate"
                hreflang="{{ .Lang }}"
                href="{{ .Permalink }}"
                />{{ end }}
  </url>
  {{ end }}
</urlset>`},
	{`_default/sitemapindex.xml`, `<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	{{ range . }}
	<sitemap>
	   	<loc>{{ .SitemapAbsURL }}</loc>
		{{ if not .LastChange.IsZero }}
	   	<lastmod>{{ .LastChange.Format "2006-01-02T15:04:05-07:00" | safeHTML }}</lastmod>
		{{ end }}
	</sitemap>
	{{ end }}
</sitemapindex>
`},
	{`disqus.html`, `{{ if .Site.DisqusShortname }}<div id="disqus_thread"></div>
<script>
    var disqus_config = function () {
    {{with .GetParam "disqus_identifier" }}this.page.identifier = '{{ . }}';{{end}}
    {{with .GetParam "disqus_title" }}this.page.title = '{{ . }}';{{end}}
    {{with .GetParam "disqus_url" }}this.page.url = '{{ . | html  }}';{{end}}
    };
    function disqusAgree(){
        localStorage.setItem("agreed_to_disqus_thread", "YES");
        localStorage.setItem("agreed_to_disqus_thread_date", (new Date()).toLocaleString() );
        location.reload();
    };
    (function() {
        if (["localhost", "127.0.0.1"].indexOf(window.location.hostname) != -1) {
            document.getElementById('disqus_thread').innerHTML = 'Disqus comments not available by default when the website is previewed locally.';
            return;
        }
        {{- if not .Site.PrivacyConfig.Disqus.SkipAgree }}
        if ((localStorage.getItem("agreed_to_disqus_thread") != "YES") ) {
          document.getElementById('disqus_thread').innerHTML =
            '{{ (default "Show comments for this page powered by [disqus.com](https://disqus.com). But first agree to the [Terms](https://help.disqus.com/terms-and-policies/terms-of-service)" (i18n "disqusTxtAgree") ) | markdownify }} '
          + '<button id="agree-to-disqus" type="button" onclick="disqusAgree()">{{default "Yes, I agree" (i18n "disqusBtnAgree")}}</button>';
          return;
        }
        {{- end }}
        var d = document, s = d.createElement('script'); s.async = true;
        s.src = '//' + {{ .Site.DisqusShortname }} + '.disqus.com/embed.js';
        s.setAttribute('data-timestamp', +new Date());
        (d.head || d.body).appendChild(s);
    })();
</script>
<noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
<a href="https://disqus.com" class="dsq-brlink">comments powered by <span class="logo-disqus">Disqus</span></a>{{end}}
`},
	{`google_analytics.html`, `{{ with .Site.GoogleAnalytics }}
<script>
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

ga('create', '{{ . }}', 'auto');
ga('send', 'pageview');
</script>
{{ end }}`},
	{`google_analytics_async.html`, `{{ with .Site.GoogleAnalytics }}
<script>
window.ga=window.ga||function(){(ga.q=ga.q||[]).push(arguments)};ga.l=+new Date;
ga('create', '{{ . }}', 'auto');
ga('send', 'pageview');
</script>
<script async src='//www.google-analytics.com/analytics.js'></script>
{{ end }}`},
	{`google_news.html`, `{{ if .IsPage }}{{ with .Params.news_keywords }}
  <meta name="news_keywords" content="{{ range $i, $kw := first 10 . }}{{ if $i }},{{ end }}{{ $kw }}{{ end }}" />
{{ end }}{{ end }}`},
	{`opengraph.html`, `<meta property="og:title" content="{{ .Title }}" />
<meta property="og:description" content="{{ with .Description }}{{ . }}{{ else }}{{if .IsPage}}{{ .Summary }}{{ else }}{{ with .Site.Params.description }}{{ . }}{{ end }}{{ end }}{{ end }}" />
<meta property="og:type" content="{{ if .IsPage }}article{{ else }}website{{ end }}" />
<meta property="og:url" content="{{ .Permalink }}" />
{{ with .Params.images }}{{ range first 6 . }}
  <meta property="og:image" content="{{ . | absURL }}" />
{{ end }}{{ end }}

{{ if .IsPage }}
{{ if not .PublishDate.IsZero }}<meta property="article:published_time" content="{{ .PublishDate.Format "2006-01-02T15:04:05-07:00" | safeHTML }}"/>
{{ else if not .Date.IsZero }}<meta property="article:published_time" content="{{ .Date.Format "2006-01-02T15:04:05-07:00" | safeHTML }}"/>{{ end }}
{{ if not .Lastmod.IsZero }}<meta property="article:modified_time" content="{{ .Lastmod.Format "2006-01-02T15:04:05-07:00" | safeHTML }}"/>{{ end }}
{{ else }}
{{ if not .Date.IsZero }}<meta property="og:updated_time" content="{{ .Date.Format "2006-01-02T15:04:05-07:00" | safeHTML }}"/>{{ end }}
{{ end }}{{ with .Params.audio }}
<meta property="og:audio" content="{{ . }}" />{{ end }}{{ with .Params.locale }}
<meta property="og:locale" content="{{ . }}" />{{ end }}{{ with .Site.Params.title }}
<meta property="og:site_name" content="{{ . }}" />{{ end }}{{ with .Params.videos }}
{{ range . }}
  <meta property="og:video" content="{{ . | absURL }}" />
{{ end }}{{ end }}

<!-- If it is part of a series, link to related articles -->
{{ $permalink := .Permalink }}
{{ $siteSeries := .Site.Taxonomies.series }}{{ with .Params.series }}
{{ range $name := . }}
  {{ $series := index $siteSeries $name }}
  {{ range $page := first 6 $series.Pages }}
    {{ if ne $page.Permalink $permalink }}<meta property="og:see_also" content="{{ $page.Permalink }}" />{{ end }}
  {{ end }}
{{ end }}{{ end }}

{{ if .IsPage }}
{{ range .Site.Authors }}{{ with .Social.facebook }}
<meta property="article:author" content="https://www.facebook.com/{{ . }}" />{{ end }}{{ with .Site.Social.facebook }}
<meta property="article:publisher" content="https://www.facebook.com/{{ . }}" />{{ end }}
<meta property="article:section" content="{{ .Section }}" />
{{ with .Params.tags }}{{ range first 6 . }}
  <meta property="article:tag" content="{{ . }}" />{{ end }}{{ end }}
{{ end }}{{ end }}

<!-- Facebook Page Admin ID for Domain Insights -->
{{ with .Site.Social.facebook_admin }}<meta property="fb:admins" content="{{ . }}" />{{ end }}`},
	{`pagination.html`, `{{ $pag := $.Paginator }}
{{ if gt $pag.TotalPages 1 }}
<ul class="pagination">
    {{ with $pag.First }}
    <li>
        <a href="{{ .URL }}" aria-label="First"><span aria-hidden="true">&laquo;&laquo;</span></a>
    </li>
    {{ end }}
    <li
    {{ if not $pag.HasPrev }}class="disabled"{{ end }}>
    <a href="{{ if $pag.HasPrev }}{{ $pag.Prev.URL }}{{ end }}" aria-label="Previous"><span aria-hidden="true">&laquo;</span></a>
    </li>
    {{ $.Scratch.Set "__paginator.ellipsed" false }}
    {{ range $pag.Pagers }}
    {{ $right := sub .TotalPages .PageNumber }}
    {{ $showNumber := or (le .PageNumber 3) (eq $right 0) }}
    {{ $showNumber := or $showNumber (and (gt .PageNumber (sub $pag.PageNumber 2)) (lt .PageNumber (add $pag.PageNumber 2)))  }}
    {{ if $showNumber }}
        {{ $.Scratch.Set "__paginator.ellipsed" false }}
        {{ $.Scratch.Set "__paginator.shouldEllipse" false }}
    {{ else }}
        {{ $.Scratch.Set "__paginator.shouldEllipse" (not ($.Scratch.Get "__paginator.ellipsed") ) }}
        {{ $.Scratch.Set "__paginator.ellipsed" true }}
    {{ end }}
    {{ if $showNumber }}
    <li
    {{ if eq . $pag }}class="active"{{ end }}><a href="{{ .URL }}">{{ .PageNumber }}</a></li>
    {{ else if ($.Scratch.Get "__paginator.shouldEllipse") }}
    <li class="disabled"><span aria-hidden="true">&hellip;</span></li>
    {{ end }}
    {{ end }}
    <li
    {{ if not $pag.HasNext }}class="disabled"{{ end }}>
    <a href="{{ if $pag.HasNext }}{{ $pag.Next.URL }}{{ end }}" aria-label="Next"><span aria-hidden="true">&raquo;</span></a>
    </li>
    {{ with $pag.Last }}
    <li>
        <a href="{{ .URL }}" aria-label="Last"><span aria-hidden="true">&raquo;&raquo;</span></a>
    </li>
    {{ end }}
</ul>
{{ end }}`},
	{`schema.html`, `{{ with .Site.Social.GooglePlus }}<link rel="publisher" href="{{ . }}"/>{{ end }}
<meta itemprop="name" content="{{ .Title }}">
<meta itemprop="description" content="{{ with .Description }}{{ . }}{{ else }}{{if .IsPage}}{{ .Summary }}{{ else }}{{ with .Site.Params.description }}{{ . }}{{ end }}{{ end }}{{ end }}">

{{if .IsPage}}{{ $ISO8601 := "2006-01-02T15:04:05-07:00" }}{{ if not .PublishDate.IsZero }}
<meta itemprop="datePublished" content="{{ .PublishDate.Format $ISO8601 | safeHTML }}" />{{ end }}
{{ if not .Date.IsZero }}<meta itemprop="dateModified" content="{{ .Date.Format $ISO8601 | safeHTML }}" />{{ end }}
<meta itemprop="wordCount" content="{{ .WordCount }}">
{{ with .Params.images }}{{ range first 6 . }}
  <meta itemprop="image" content="{{ . | absURL }}">
{{ end }}{{ end }}

<!-- Output all taxonomies as schema.org keywords -->
<meta itemprop="keywords" content="{{ if .IsPage}}{{ range $index, $tag := .Params.tags }}{{ $tag }},{{ end }}{{ else }}{{ range $plural, $terms := .Site.Taxonomies }}{{ range $term, $val := $terms }}{{ printf "%s," $term }}{{ end }}{{ end }}{{ end }}" />
{{ end }}`},
	{`shortcodes/figure.html`, `<!-- image -->
<figure{{ with .Get "class" }} class="{{.}}"{{ end }}>
    {{ if .Get "link"}}<a href="{{ .Get "link" }}"{{ with .Get "target" }} target="{{ . }}"{{ end }}{{ with .Get "rel" }} rel="{{ . }}"{{ end }}>{{ end }}
        <img src="{{ .Get "src" }}" {{ if or (.Get "alt") (.Get "caption") }}alt="{{ with .Get "alt"}}{{.}}{{else}}{{ .Get "caption" }}{{ end }}" {{ end }}{{ with .Get "width" }}width="{{.}}" {{ end }}{{ with .Get "height" }}height="{{.}}" {{ end }}/>
    {{ if .Get "link"}}</a>{{ end }}
    {{ if or (or (.Get "title") (.Get "caption")) (.Get "attr")}}
    <figcaption>{{ if isset .Params "title" }}
        <h4>{{ .Get "title" }}</h4>{{ end }}
        {{ if or (.Get "caption") (.Get "attr")}}<p>
        {{ .Get "caption" }}
        {{ with .Get "attrlink"}}<a href="{{.}}"> {{ end }}
            {{ .Get "attr" }}
        {{ if .Get "attrlink"}}</a> {{ end }}
        </p> {{ end }}
    </figcaption>
    {{ end }}
</figure>
<!-- image -->`},
	{`shortcodes/gist.html`, `<script src="//gist.github.com/{{ index .Params 0 }}/{{ index .Params 1 }}.js{{if len .Params | eq 3 }}?file={{ index .Params 2 }}{{end}}"></script>`},
	{`shortcodes/highlight.html`, `{{ if len .Params | eq 2 }}{{ highlight (trim .Inner "\n\r") (.Get 0) (.Get 1) }}{{ else }}{{ highlight (trim .Inner "\n\r") (.Get 0) "" }}{{ end }}`},
	{`shortcodes/instagram.html`, `{{ if len .Params | eq 2 }}{{ if eq (.Get 1) "hidecaption" }}{{ with getJSON "https://api.instagram.com/oembed/?url=https://instagram.com/p/" (index .Params 0) "/&hidecaption=1" }}{{ .html | safeHTML }}{{ end }}{{ end }}{{ else }}{{ with getJSON "https://api.instagram.com/oembed/?url=https://instagram.com/p/" (index .Params 0) "/&hidecaption=0" }}{{ .html | safeHTML }}{{ end }}{{ end }}`},
	{`shortcodes/ref.html`, `{{ if len .Params | eq 2 }}{{ ref .Page (.Get 0) (.Get 1) }}{{ else }}{{ ref .Page (.Get 0) }}{{ end }}`},
	{`shortcodes/relref.html`, `{{ if len .Params | eq 2 }}{{ relref .Page (.Get 0) (.Get 1) }}{{ else }}{{ relref .Page (.Get 0) }}{{ end }}`},
	{`shortcodes/speakerdeck.html`, `<script async class='speakerdeck-embed' data-id='{{ index .Params 0 }}' data-ratio='1.33333333333333' src='//speakerdeck.com/assets/embed.js'></script>`},
	{`shortcodes/tweet.html`, `{{ (getJSON "https://api.twitter.com/1/statuses/oembed.json?id=" (index .Params 0)).html | safeHTML }}`},
	{`shortcodes/vimeo.html`, `{{ if .IsNamedParams }}<div {{ if .Get "class" }}class="{{ .Get "class" }}"{{ else }}style="position: relative; padding-bottom: 56.25%; padding-top: 30px; height: 0; overflow: hidden;"{{ end }}>
  <iframe src="//player.vimeo.com/video/{{ .Get "id" }}" {{ if not (.Get "class") }}style="position: absolute; top: 0; left: 0; width: 100%; height: 100%;" {{ end }}webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
 </div>{{ else }}
<div {{ if len .Params | eq 2 }}class="{{ .Get 1 }}"{{ else }}style="position: relative; padding-bottom: 56.25%; padding-top: 30px; height: 0; overflow: hidden;"{{ end }}>
  <iframe src="//player.vimeo.com/video/{{ .Get 0 }}" {{ if len .Params | eq 1 }}style="position: absolute; top: 0; left: 0; width: 100%; height: 100%;" {{ end }}webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
 </div>
{{ end }}`},
	{`shortcodes/youtube.html`, `{{ if .IsNamedParams }}
<div {{ if .Get "class" }}class="{{ .Get "class" }}"{{ else }}style="position: relative; padding-bottom: 56.25%; padding-top: 30px; height: 0; overflow: hidden;"{{ end }}>
  <iframe src="//www.youtube.com/embed/{{ .Get "id" }}?{{ with .Get "autoplay" }}{{ if eq . "true" }}autoplay=1{{ end }}{{ end }}"
  {{ if not (.Get "class") }}style="position: absolute; top: 0; left: 0; width: 100%; height: 100%;" {{ end }}allowfullscreen frameborder="0" title="YouTube Video"></iframe>
</div>{{ else }}
<div {{ if len .Params | eq 2 }}class="{{ .Get 1 }}"{{ else }}style="position: relative; padding-bottom: 56.25%; padding-top: 30px; height: 0; overflow: hidden;"{{ end }}>
  <iframe src="//www.youtube.com/embed/{{ .Get 0 }}" {{ if len .Params | eq 1 }}style="position: absolute; top: 0; left: 0; width: 100%; height: 100%;" {{ end }}allowfullscreen frameborder="0" title="YouTube Video"></iframe>
 </div>
{{ end }}`},
	{`twitter_cards.html`, `{{- with $.Params.images -}}
<meta name="twitter:card" content="summary_large_image"/>
<meta name="twitter:image" content="{{ index . 0 | absURL }}"/>
{{ else -}}
{{- $images := $.Resources.ByType "image" -}}
{{- $featured := $images.GetMatch "*feature*" -}}
{{- $featured := cond (ne $featured nil) $featured ($images.GetMatch "{*cover*,*thumbnail*}") -}}
{{- with $featured -}}
<meta name="twitter:card" content="summary_large_image"/>
<meta name="twitter:image" content="{{ $featured.Permalink }}"/>
{{- else -}}
{{- with $.Site.Params.images -}}
<meta name="twitter:card" content="summary_large_image"/>
<meta name="twitter:image" content="{{ index . 0 | absURL }}"/>
{{ else -}}
<meta name="twitter:card" content="summary"/>
{{- end -}}
{{- end -}}
{{- end }}
<meta name="twitter:title" content="{{ .Title }}"/>
<meta name="twitter:description" content="{{ with .Description }}{{ . }}{{ else }}{{if .IsPage}}{{ .Summary }}{{ else }}{{ with .Site.Params.description }}{{ . }}{{ end }}{{ end }}{{ end -}}"/>
{{ with .Site.Social.twitter -}}
<meta name="twitter:site" content="@{{ . }}"/>
{{ end -}}
{{ range .Site.Authors }}
{{ with .twitter -}}
<meta name="twitter:creator" content="@{{ . }}"/>
{{ end -}}
{{ end -}}`},
}
