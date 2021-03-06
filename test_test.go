package main

import (
	"fmt"
	"net/url"
	"regexp"
	"testing"
)

var str = `
<?xml version="1.0" encoding="UTF-8"?><rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:anchor="https://anchor.fm/xmlns">
	<channel>
		<title><![CDATA[【睡前相声】郭德纲于谦相声合集-高清晰]]></title>
		<description><![CDATA[【睡前相声】郭德纲于谦相声合集-高清晰 #郭德纲 #于谦 #相声 #高清 #德云社

此频道没有任何收入，个人兴趣，由视频转音频之后上传发布。

如有任何问题，请联系：lionft@qq.com]]></description>
		<link>https://iiba.fun</link>
		<image>
			<url>https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg</url>
			<title>【睡前相声】郭德纲于谦相声合集-高清晰</title>
			<link>https://iiba.fun</link>
		</image>
		<generator>Anchor Podcasts</generator>
		<lastBuildDate>Tue, 14 Jun 2022 11:56:17 GMT</lastBuildDate>
		<atom:link href="https://anchor.fm/s/590ba140/podcast/rss" rel="self" type="application/rss+xml"/>
		<author><![CDATA[D.M.]]></author>
		<copyright><![CDATA[D.M.]]></copyright>
		<language><![CDATA[zh-cn]]></language>
		<atom:link rel="hub" href="https://pubsubhubbub.appspot.com/"/>
		<itunes:author>D.M.</itunes:author>
		<itunes:summary>【睡前相声】郭德纲于谦相声合集-高清晰 #郭德纲 #于谦 #相声 #高清 #德云社

此频道没有任何收入，个人兴趣，由视频转音频之后上传发布。

如有任何问题，请联系：lionft@qq.com</itunes:summary>
		<itunes:type>episodic</itunes:type>
		<itunes:owner>
			<itunes:name>D.M.</itunes:name>
			<itunes:email>podcasts60+590ba140@anchor.fm</itunes:email>
		</itunes:owner>
		<itunes:explicit>No</itunes:explicit>
		<itunes:category text="Arts"/>
		<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
		<item>
			<title><![CDATA[【相声音频助眠】《谁叫于谦》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《谁叫于谦》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1ft1k9</link>
			<guid isPermaLink="false">7dafd69b-d658-4c37-ba17-46747e3e6fdc</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 18 Mar 2022 08:53:40 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/49235017/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-18%2Fbf950af3-3925-d32a-20ca-c0304419276f.mp3" length="51740863" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《谁叫于谦》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3228</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《超级玛丽》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《超级玛丽》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1ft1jt</link>
			<guid isPermaLink="false">b60c2621-fea3-4835-89ee-ab8c4fe6bf19</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 18 Mar 2022 08:53:10 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/49235005/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-18%2F1e4f8191-e36e-2d3e-757a-6546da0565ff.mp3" length="46058018" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《超级玛丽》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2872</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《摸王八》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《摸王八》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1ft1jn</link>
			<guid isPermaLink="false">ce3335a6-b39c-49cf-b904-bc41a4cef232</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 18 Mar 2022 08:52:43 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/49234999/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-18%2F3e3a8003-58e5-5699-b584-d76a6133ff0d.mp3" length="62271097" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《摸王八》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3886</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《你来一趟》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《你来一趟》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1ft1in</link>
			<guid isPermaLink="false">ab5a4d66-e085-4ff7-9d82-5f9e01da6865</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 18 Mar 2022 08:52:02 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/49234967/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-18%2F06733bd4-f045-50ac-4d27-4536b9392c23.mp3" length="54819396" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《你来一趟》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3420</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《来一颗吗》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《来一颗吗》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f92kr</link>
			<guid isPermaLink="false">d1473f2c-389a-4bf0-a83f-6f56808acbb8</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 05 Mar 2022 06:55:52 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48580699/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-5%2F35bc7df1-8d46-b038-0078-f097b90e7853.mp3" length="60778692" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《来一颗吗》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3794</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《大伙真棒》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《大伙真棒》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f92km</link>
			<guid isPermaLink="false">94f39429-cb6c-4856-8151-9b7ccc3902df</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 05 Mar 2022 06:55:33 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48580694/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-5%2F5dd0431c-ae0d-ff6d-88f6-0fde57d4e1cd.mp3" length="67010935" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《大伙真棒》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4183</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《男一号裸替》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《男一号裸替》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f92kd</link>
			<guid isPermaLink="false">91015895-913c-443b-99f2-0a3320d78447</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 05 Mar 2022 06:55:09 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48580685/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-5%2Fc76c683a-91ad-3292-16ab-12c769875776.mp3" length="56745493" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《男一号裸替》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3541</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《大姨夫来了》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《大姨夫来了》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f92it</link>
			<guid isPermaLink="false">56333fae-d871-4a40-8c42-db1009c7d382</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 05 Mar 2022 06:51:57 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48580637/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-5%2F17361c59-f907-dea0-05d1-7bb9bd53b284.mp3" length="58749031" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《大姨夫来了》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3667</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《人无完人》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《人无完人》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f92in</link>
			<guid isPermaLink="false">371a8c6e-89bb-44ac-a1b4-db2f3880d45b</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 05 Mar 2022 06:51:38 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48580631/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-5%2Fd8a0dac8-88b4-cdb1-c654-48d711efa8a3.mp3" length="76796738" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《人无完人》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4795</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《鸡贼王中王》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《鸡贼王中王》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2pmd</link>
			<guid isPermaLink="false">95fcb709-7cb4-4c42-81a5-2b6014a554db</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:44:53 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374925/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2F2d8dbfd7-c318-bf11-dc6d-35597fd2b405.mp3" length="58181796" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《鸡贼王中王》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3631</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《老大个了》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《老大个了》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2pm2</link>
			<guid isPermaLink="false">b76e6ba1-23ad-427b-880d-adb10ca6d0a7</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:44:30 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374914/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2Fef446faf-8aa8-f805-6533-a558c9abf4e8.mp3" length="83625412" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《老大个了》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>5222</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《三天假期》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《三天假期》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2plt</link>
			<guid isPermaLink="false">639addcd-a2f2-423a-abc7-81370f0f7f69</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:44:06 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374909/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2F6d677698-4663-4ffa-e467-a6225031324b.mp3" length="65654239" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《三天假期》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4098</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《吃喝拉撒》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《吃喝拉撒》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2plo</link>
			<guid isPermaLink="false">90c60a34-c6cd-4ab7-b227-35063e6fff2a</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:43:45 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374904/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2F3c21021e-a39f-f7ac-67fa-91756383927f.mp3" length="44378141" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《吃喝拉撒》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2768</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《瞄准儿呢》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《瞄准儿呢》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2plj</link>
			<guid isPermaLink="false">beff3cd3-c1f5-4ab5-bf87-986e9098805f</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:43:26 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374899/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2Fa39e8555-8783-ecd6-1af4-5d814db9dac0.mp3" length="71800651" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《瞄准儿呢》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4482</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《诶》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《诶》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2pld</link>
			<guid isPermaLink="false">e153b7ba-21d8-4055-9217-273417193b8f</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:43:08 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374893/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2F2c026d0d-8247-9a2f-6d8a-76581bdb32d0.mp3" length="71481095" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《诶》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4463</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《远古时期》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《远古时期》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2pl3</link>
			<guid isPermaLink="false">7374004a-eb5b-4916-9471-cf018884eabc</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:42:52 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374883/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2Fed051b12-083f-c0c6-9a42-25514d85093f.mp3" length="78331499" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《远古时期》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4891</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《于家大菊花》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《于家大菊花》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2pkv</link>
			<guid isPermaLink="false">902edb3a-79c9-41fc-8bf1-058230b404c1</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:42:34 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374879/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2F180b1cf9-8ef4-0ffb-24ca-6c3a66f223de.mp3" length="57504811" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《于家大菊花》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3589</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《殿下》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《殿下》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2pkn</link>
			<guid isPermaLink="false">96e20578-a376-4113-b2b1-06ebb291c774</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:42:10 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374871/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2F303515b3-4d5a-afc1-4b41-29645132861e.mp3" length="64838104" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《殿下》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4048</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《四合院儿》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《四合院儿》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2pkj</link>
			<guid isPermaLink="false">69338aab-4def-44e9-bbe8-eefc28172f40</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:41:52 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374867/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2F45e00bad-02fc-91ac-3157-d81074386457.mp3" length="58179791" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《四合院儿》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3631</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《一视同仁》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《一视同仁》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2pkc</link>
			<guid isPermaLink="false">4ce7c043-2901-4bd2-b330-41b3d5314df3</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:41:35 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374860/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2F44d852e8-4053-c487-6087-cee1c2cf2f49.mp3" length="49795365" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《一视同仁》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3107</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《他是亲的》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《他是亲的》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2pk7</link>
			<guid isPermaLink="false">0bae675b-9a0d-40ed-a05d-091d6bac119e</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:41:11 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374855/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2Fea53cbff-4f25-3f0f-ae6c-adfc3d7eb42b.mp3" length="53204317" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《他是亲的》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3320</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《以合为贵》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《以合为贵》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f2pjv</link>
			<guid isPermaLink="false">74074990-7093-49cd-93d7-a640f4bb1a55</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Tue, 01 Mar 2022 07:40:48 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48374847/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-1%2Fbd9d8006-f958-0bdf-680a-b46bdb459cd2.mp3" length="66627839" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《以合为贵》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4159</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《冰清玉洁》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《冰清玉洁》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0r38</link>
			<guid isPermaLink="false">6ac0a78f-45ca-4b8b-93b7-856bd3eaebab</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 03:49:45 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48310824/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2Fd6aab6fb-7e80-5159-e0f3-aebf256874c6.mp3" length="55028059" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《冰清玉洁》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3434</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《偷儿胡同》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《偷儿胡同》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0qts</link>
			<guid isPermaLink="false">a23758f1-4699-4c8e-812e-6dcc3edf9572</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 03:45:09 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48310652/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2F5979d236-b50c-7777-47b4-02e5be00b2a0.mp3" length="64280197" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《偷儿胡同》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4012</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《两位公公》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《两位公公》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0qte</link>
			<guid isPermaLink="false">0a4bf35a-6dc1-41b9-83ab-cf4583e4f09b</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 03:44:47 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48310638/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2F33f39905-957d-3fe7-55cc-b8fa816a7cf5.mp3" length="60523938" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《两位公公》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3778</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《于十娘》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《于十娘》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0qst</link>
			<guid isPermaLink="false">4af658f5-be48-45f4-af9e-b2c50b17f360</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 03:44:22 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48310621/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2F59df19b3-7e4a-dcbb-d3d1-0fbb2c61fe84.mp3" length="75140949" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《于十娘》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4691</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《配合一下》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《配合一下》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0qsi</link>
			<guid isPermaLink="false">dc44daef-2f78-4b0f-9989-532617781e89</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 03:43:59 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48310610/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2Fbf124fe9-727b-a9b0-962e-286bce153640.mp3" length="55058825" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《配合一下》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3436</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《大爷大爷》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《大爷大爷》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0qgo</link>
			<guid isPermaLink="false">338ca002-df55-4114-a61d-a476dab7dd73</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 03:38:53 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48310232/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2Ffc58794c-e1cd-254f-90e5-e683723e7b34.mp3" length="71795812" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《大爷大爷》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4482</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《二怎么写》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《二怎么写》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0qgi</link>
			<guid isPermaLink="false">30f74d50-947b-43b6-b5c8-8c2152abd08e</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 03:38:33 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48310226/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2Ff177f584-db09-0f4a-5e73-cccb284b1ed7.mp3" length="68589033" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《二怎么写》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4282</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《冬天游泳》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《冬天游泳》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0q71</link>
			<guid isPermaLink="false">268dbc6c-7a6c-4525-be2b-84abafdf25b4</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 03:29:51 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48309921/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2F173794c4-498e-c8bc-dcf7-343d7c6bfad4.mp3" length="50798295" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《冬天游泳》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3170</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《女足》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《女足》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0q6a</link>
			<guid isPermaLink="false">77a840c2-7afe-4c22-8297-7b63496e16dc</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 03:29:29 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48309898/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2Fe078d8f0-63e5-2b3e-439a-e06d640df977.mp3" length="21535232" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《女足》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1342</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《冬奥会》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《冬奥会》</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0q5u</link>
			<guid isPermaLink="false">145afa95-7a32-4095-9162-86784d2a50ce</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 03:28:58 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48309886/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2Faaa325a7-20b3-5b5b-ccbd-c04d7ca0617d.mp3" length="60648242" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《冬奥会》&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3786</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《你来啦》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《你来啦》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0nki</link>
			<guid isPermaLink="false">9981f14f-b852-42b2-b6d1-a3c39494ad14</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 02:24:41 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48307282/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2F1438842d-2404-2062-5540-13f740c68225.mp3" length="49063767" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《你来啦》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3061</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《我很厚道》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《我很厚道》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0njp</link>
			<guid isPermaLink="false">01a6ab3b-5bdf-48ed-8ae3-04c331bc7f24</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 02:24:13 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48307257/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2F6bedeffa-02d9-982b-44d1-3b85db06d2f5.mp3" length="44229541" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《我很厚道》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2759</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《冲啊冲啊》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《冲啊冲啊》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0niq</link>
			<guid isPermaLink="false">72362e9b-2df3-49a2-8dfe-2b1f586e5d20</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 02:23:44 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48307226/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2Ff1fb4a68-d864-256b-648e-fccdbfa91080.mp3" length="54697128" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《冲啊冲啊》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3413</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《新年快乐》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《新年快乐》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0ni9</link>
			<guid isPermaLink="false">8bfdf72f-537c-4608-8f81-59a0d62dc33b</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 02:23:09 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48307209/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2Fe08e8222-bd8a-34dc-f9a9-f740810bdcdd.mp3" length="72486884" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《新年快乐》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4525</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《退票》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《退票》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0nhs</link>
			<guid isPermaLink="false">560b103a-5c5f-4c9c-bca4-d3ec903b6df6</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 02:22:40 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48307196/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2Fbc6a927b-78f2-8a2a-4dfd-b6ea969a23dc.mp3" length="52422803" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《退票》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3272</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《包子铺》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《包子铺》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0nfh</link>
			<guid isPermaLink="false">3163f42a-1cfd-4180-9b56-d4344d12d6ba</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 02:21:14 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48307121/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2F2cf523bd-e663-1d82-1129-1dc876dcb0f6.mp3" length="45949937" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《包子铺》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2867</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《火车爆胎了》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《火车爆胎了》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0n9u</link>
			<guid isPermaLink="false">2b1876ab-7206-487d-ad25-c40771881819</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 02:15:14 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48306942/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2F0a48e5d6-5bfe-cf1f-b720-c6e14a1c2ff3.mp3" length="71410390" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《火车爆胎了》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4458</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《你跟着我》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《你跟着我》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0n9f</link>
			<guid isPermaLink="false">ac70759f-2375-4492-a9d4-697d47293bd5</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 02:14:43 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48306927/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2F1b065a52-6b1c-fd8a-eac3-0127ba8aa710.mp3" length="54222985" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《你跟着我》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3384</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《摄影的祖先》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《摄影的祖先》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0n7s</link>
			<guid isPermaLink="false">7a8639d8-b666-458c-b8cb-98eab03d3fa5</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 02:14:05 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48306876/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2Fb1f27492-24ec-ab78-eadd-f465bbb88f2e.mp3" length="59813367" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《摄影的祖先》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3733</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《猫咪》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《猫咪》<strong>#郭德纲 #于谦 #德云社 #相声</strong></p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1f0n72</link>
			<guid isPermaLink="false">0b0ddd8e-5b00-416d-9a37-0b4596c34eb5</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Mon, 28 Feb 2022 02:12:22 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48306850/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-28%2F08476336-5215-a327-28ac-b8da97769817.mp3" length="56749341" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《猫咪》&lt;strong&gt;#郭德纲 #于谦 #德云社 #相声&lt;/strong&gt;&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3542</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《等会儿等会儿》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《等会儿等会儿》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1evhoe</link>
			<guid isPermaLink="false">c5f193bc-3594-4e62-80bf-37c9292555cb</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 27 Feb 2022 04:43:49 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48268494/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-27%2Fb0d28dcb-1c65-4616-dbdf-1ea48fddc787.mp3" length="55103711" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《等会儿等会儿》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3439</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《强行慰问演出》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《强行慰问演出》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1evho8</link>
			<guid isPermaLink="false">56d5e840-e70d-431c-afde-36778a554896</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 27 Feb 2022 04:43:16 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48268488/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-27%2F06e96abd-15cb-1aab-df19-7ee6a61d2b55.mp3" length="62212498" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《强行慰问演出》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3883</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《寒窑赋》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《寒窑赋》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1evhnm</link>
			<guid isPermaLink="false">7951f140-e619-435c-94db-af91787be0b0</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 27 Feb 2022 04:42:48 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48268470/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-27%2Fb8a6454e-8b25-7bd9-f861-7f4f25751544.mp3" length="76798282" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《寒窑赋》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4795</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《青梅竹马》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《青梅竹马》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1evhn8</link>
			<guid isPermaLink="false">ace01f6a-33f9-43f6-9bd7-ec159b57f05b</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 27 Feb 2022 04:42:20 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48268456/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-27%2F07c041f4-eb2c-52e1-f1f4-34cfb32eca62.mp3" length="52763948" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《青梅竹马》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3293</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《非礼树》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《非礼树》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1evhmp</link>
			<guid isPermaLink="false">1d3ffbb8-33c3-4df7-959f-5f970923d7a8</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 27 Feb 2022 04:41:46 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48268441/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-27%2F4dedb8eb-f38b-aee9-04aa-527d0e91d6e9.mp3" length="59778968" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《非礼树》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3731</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《我要当关公》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《我要当关公》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1evhmc</link>
			<guid isPermaLink="false">a8bdee3b-4a95-4dd9-9aac-2752ac601371</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 27 Feb 2022 04:41:11 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48268428/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-27%2Fcf55dbfa-6bbc-f507-5290-09cdf55591f3.mp3" length="52140859" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《我要当关公》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3253</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《这事儿闹的》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《这事儿闹的》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1evhm5</link>
			<guid isPermaLink="false">b1cb5da0-60b4-403d-a2a6-7216c8bde5bc</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 27 Feb 2022 04:40:37 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48268421/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-27%2F41dc335e-18e8-f635-044f-67e0ce36df7c.mp3" length="58531169" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《这事儿闹的》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3653</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《你要蛋定》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《你要蛋定》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1evhlr</link>
			<guid isPermaLink="false">a83e0614-0d5b-474f-9bb6-b05474516ea9</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 27 Feb 2022 04:40:07 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48268411/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-27%2F61c7a5a4-7340-9a9d-b9fa-4b1ae05e65f0.mp3" length="65053234" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《你要蛋定》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4061</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《霍喽霍喽》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《霍喽霍喽》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1evhl2</link>
			<guid isPermaLink="false">cd429866-52f2-4008-a752-d033156f8488</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 27 Feb 2022 04:39:22 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48268386/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-27%2F2dc1ca3a-9919-eb0a-28d8-9ff35c101a19.mp3" length="62247701" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《霍喽霍喽》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3885</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《一边儿玩去》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《一边儿玩去》#郭德纲 #于谦 #德云社 #相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1evhk5</link>
			<guid isPermaLink="false">25ea70c4-1d4b-4485-91c0-9546eef82ea6</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 27 Feb 2022 04:38:17 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48268357/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-27%2F259d7f8b-8397-2c34-5aa4-6a4173d4f78d.mp3" length="55102939" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《一边儿玩去》#郭德纲 #于谦 #德云社 #相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3439</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《一扎来长》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《一扎来长》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1er5dh</link>
			<guid isPermaLink="false">8a52c702-6f9e-46ce-b73d-a899dcdeab34</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 24 Feb 2022 01:13:31 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48124785/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-24%2Fe6a72b6a-611c-f211-af04-86d428b88315.mp3" length="70856902" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《一扎来长》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4423</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《咱这讲理》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《咱这讲理》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1er5ch</link>
			<guid isPermaLink="false">4b8985dc-53ed-4497-806d-78d733f51ed2</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 24 Feb 2022 01:13:00 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48124753/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-23%2F57b1b291-50b0-33de-59c3-7e322b4373de.mp3" length="31242262" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《咱这讲理》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1947</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《早说啊》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《早说啊》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1er5c4</link>
			<guid isPermaLink="false">7a817ff8-7eaa-4b3e-98aa-1f804dc0b0cb</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 24 Feb 2022 01:12:12 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48124740/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-23%2Fbb936660-98c2-0c20-722e-766f8d7a1851.mp3" length="55158644" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《早说啊》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3442</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《往这拍一下》]]></title>
			<description><![CDATA[<p><strong>【相声音频助眠】《往这拍一下》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社</strong></p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1eptb6</link>
			<guid isPermaLink="false">2cb2378f-8ff9-43bd-a572-430155c4087a</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 23 Feb 2022 10:04:58 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48083750/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-23%2Fef5b38db-d8e0-a2b1-f828-57de5623a81d.mp3" length="48773969" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;&lt;strong&gt;【相声音频助眠】《往这拍一下》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社&lt;/strong&gt;&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3043</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《风卷残云》]]></title>
			<description><![CDATA[<p><strong>【相声音频助眠】《风卷残云》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社</strong></p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1epta0</link>
			<guid isPermaLink="false">3f1121fb-3137-4799-8d32-1f7ec28025f2</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 23 Feb 2022 09:20:42 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48083712/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-23%2F417e15f7-edbc-e1e9-5afe-0ded8f8c8bf5.mp3" length="58235397" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;&lt;strong&gt;【相声音频助眠】《风卷残云》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社&lt;/strong&gt;&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3634</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《嗯嗯嗯嗯嗯》]]></title>
			<description><![CDATA[<p><strong>【相声音频助眠】《嗯嗯嗯嗯嗯》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社</strong></p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1ept9i</link>
			<guid isPermaLink="false">8ad2fd3c-146e-4f29-91df-0227307ae898</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 23 Feb 2022 09:19:58 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48083698/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-23%2F366abb96-97cf-e991-761d-e0d273e29f0d.mp3" length="65722645" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;&lt;strong&gt;【相声音频助眠】《嗯嗯嗯嗯嗯》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社&lt;/strong&gt;&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4102</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《白天捏～》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《白天捏～》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1epr2s</link>
			<guid isPermaLink="false">aee119c7-440e-4b4b-8909-9daf9cedb668</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 23 Feb 2022 07:44:01 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48081436/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-23%2Fdaa77ebc-674e-fc0c-7a0e-0b8516545ddc.mp3" length="68826680" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《白天捏～》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4297</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《特别可爱》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《特别可爱》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1epr24</link>
			<guid isPermaLink="false">5f9601a7-2fdd-4af9-b3fb-e096d3233a37</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 23 Feb 2022 07:42:44 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48081412/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-23%2F1d76d94c-6ef3-18d7-bb13-937dbb076ee7.mp3" length="57574317" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《特别可爱》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3593</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【相声音频助眠】《情人》]]></title>
			<description><![CDATA[<p>【相声音频助眠】《情人》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1epr0n</link>
			<guid isPermaLink="false">cfbdffa4-843a-45c7-9046-e9560f4fa285</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 23 Feb 2022 07:41:14 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/48081367/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-1-23%2Ffdade908-a210-a283-eb52-3f1215cab3d4.mp3" length="51553019" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【相声音频助眠】《情人》#2022德云社 #无损音质 #郭德纲 #于谦 #德云社&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3217</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《呦呵呦》无损音质]]></title>
			<description><![CDATA[<p>2021德云社郭德纲于谦《呦呵呦》无损音质</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2hq</link>
			<guid isPermaLink="false">a6ab51fc-0d54-4f99-9dfc-e7335c63b104</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:19:00 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422522/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-7-7%2F26d0e1b9-5d9c-d63a-5f0b-5f9e0c7f5d01.mp3" length="44942632" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;2021德云社郭德纲于谦《呦呵呦》无损音质&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2804</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《这里真热闹》]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《这里真热闹》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2hh</link>
			<guid isPermaLink="false">af4848bd-c5bd-43b1-915e-4f792171f9d4</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:17:22 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422513/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F32b4e236-8c8e-47f8-090a-6341d2318aa3.mp3" length="23690034" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《这里真热闹》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1469</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《我要买房》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《我要买房》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2h8</link>
			<guid isPermaLink="false">a322440d-0137-4bfc-bad4-43570b4f9929</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:16:58 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422504/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2Fa00de589-7ae3-af6c-cc6a-a24328d712be.mp3" length="57360219" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《我要买房》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3577</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《学五行诗》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《学五行诗》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2h3</link>
			<guid isPermaLink="false">a2fe988c-87c3-403f-a033-8f5cf97b5010</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:16:33 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422499/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F9bc47143-fb4e-f1a0-b059-9c690f2a7ade.mp3" length="36289389" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《学五行诗》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2259</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《把家长找来》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《把家长找来》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2gp</link>
			<guid isPermaLink="false">689373f5-2551-4f25-8987-f4be5f30a8d9</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:16:07 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422489/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F25ba038a-3ed3-d736-850b-24ee61c21972.mp3" length="33582678" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《把家长找来》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2091</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《擦玻璃》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《擦玻璃》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2gn</link>
			<guid isPermaLink="false">4abb20c3-b21d-4556-9838-98caddf4e15d</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:15:46 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422487/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F79dae2d0-18dd-c31e-d88e-4766e980c906.mp3" length="30979905" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《擦玻璃》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1928</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《武林中人》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《武林中人》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2gf</link>
			<guid isPermaLink="false">a48fd283-6b60-42c3-a18d-8e11d3ff04c5</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:15:23 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422479/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2Fcb0cc482-49c1-5520-9a81-9b757681dd5e.mp3" length="32694302" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《武林中人》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2036</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《全世界走走》高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《全世界走走》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2fu</link>
			<guid isPermaLink="false">244f2686-8629-4a2c-afb8-b6474f549761</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:14:53 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422462/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F36c70ae1-1ad1-91b6-9842-35d21e6ecf60.mp3" length="33644518" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《全世界走走》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2095</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《老冒儿》高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《老冒儿》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2fr</link>
			<guid isPermaLink="false">2973fb48-4632-4113-9d02-1bba8b9f5575</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:14:28 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422459/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F1f9085a0-1051-2701-4d27-de522040561c.mp3" length="28268991" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《老冒儿》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1759</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《郭老大开饭店》高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《郭老大开饭店》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2fj</link>
			<guid isPermaLink="false">a8f972d6-c0d6-4bed-87b3-a81ecc5c8c76</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:14:05 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422451/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F169a3466-3e8a-3ee6-80d1-17633f2eb5a5.mp3" length="24029222" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《郭老大开饭店》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1494</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《出类拔萃》高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《出类拔萃》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2f7</link>
			<guid isPermaLink="false">5dfd515c-f6ed-47bc-b1a1-df8a972fad22</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:13:22 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422439/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2Feabd51de-959f-d74e-0d53-6ecca5ee2da5.mp3" length="27006341" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《出类拔萃》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1680</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《八盼》高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《八盼》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2ev</link>
			<guid isPermaLink="false">f51c02a3-79ca-4302-ae01-7aba783a0cb1</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:12:56 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422431/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F85e86f92-fce8-e891-58fc-406e69a57674.mp3" length="32629688" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《八盼》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2032</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《是秃子总会发光》高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《是秃子总会发光》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2ei</link>
			<guid isPermaLink="false">a6142d2f-afdc-43c3-83e9-1e9872de2364</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:12:13 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422418/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F2d46f8f4-9a16-7d4f-3a65-ac2b0589fd36.mp3" length="24281858" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《是秃子总会发光》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1510</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《貂皮裤衩》高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《貂皮裤衩》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2ea</link>
			<guid isPermaLink="false">e860c6f8-53ff-48f9-b9fe-1ab83f06030c</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:11:39 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422410/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F5f0e0524-6280-2e67-bf2a-46e00dfd434f.mp3" length="29785588" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《貂皮裤衩》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1854</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《成功人士》高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《成功人士》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2e0</link>
			<guid isPermaLink="false">d28f0723-84fc-43df-9a02-ba63847ad54c</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:11:09 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422400/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2Fdf51f6fd-c072-8f51-3aff-3f9e16813ea1.mp3" length="32080395" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《成功人士》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1997</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《佛在心头》高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《佛在心头》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2dn</link>
			<guid isPermaLink="false">0799830d-5435-4d57-917f-d2089a105ff1</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:10:43 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422391/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F8febd00d-f44e-3701-f679-c75ae6457680.mp3" length="25938349" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《佛在心头》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1613</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《语言的艺术》高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《语言的艺术》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2d7</link>
			<guid isPermaLink="false">a5617492-4234-48b7-9906-44698cb1b29f</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:10:14 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422375/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F3dc6cda5-6c81-52fb-b046-4f4a5f809ee0.mp3" length="26757515" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《语言的艺术》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1665</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《于家爷俩非洲游》高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《于家爷俩非洲游》高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e15j2ce</link>
			<guid isPermaLink="false">6913e326-c390-4999-8741-f9bcc0cfc98b</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sat, 07 Aug 2021 09:08:25 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/38422350/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F4c0627f7-d8ac-80db-88d9-c3cf9aacacec.mp3" length="8609382" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《于家爷俩非洲游》高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>529</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《论朋友》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《论朋友》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e13t4if</link>
			<guid isPermaLink="false">45af2aed-41f6-4cbb-a78d-fc38062357f8</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 04 Jul 2021 07:43:08 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/36655119/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2Faa90a8f6-6eda-65d7-d800-1e1f1056db12.mp3" length="32510192" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《论朋友》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2024</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《比您差点》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《比您差点》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e13t4i6</link>
			<guid isPermaLink="false">fe74239f-5d05-4701-85fb-b88265b63c3e</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 04 Jul 2021 07:42:47 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/36655110/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2Fbebe70eb-b836-a14a-729c-2a48f82e8e5d.mp3" length="29598071" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《比您差点》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1842</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《你要坚强》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《你要坚强》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e13t4i1</link>
			<guid isPermaLink="false">31f49052-d3cb-4c3f-895f-2c3bb8416e21</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 04 Jul 2021 07:42:25 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/36655105/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2Fce5230fb-5dc6-4b34-7a78-33f9c0a71aa8.mp3" length="32206325" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《你要坚强》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2005</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【特别放送】《十八摸》 高清音频 唱已调]]></title>
			<description><![CDATA[<p>【特别放送】郭德纲于谦相声《十八摸》 高清音频 唱已调</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e13t4gt</link>
			<guid isPermaLink="false">5f1c8003-c265-4814-a73e-e5f54c58ec4b</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 04 Jul 2021 07:40:56 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/36655069/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2Fe0bce92e-36cd-8c74-5833-931343bf72b2.mp3" length="39866101" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【特别放送】郭德纲于谦相声《十八摸》 高清音频 唱已调&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2484</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《一夫一妻制》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《一夫一妻制》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e13t4gm</link>
			<guid isPermaLink="false">c1b18ff4-154b-4239-b903-441d6b9e8f57</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 04 Jul 2021 07:40:34 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/36655062/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F52c7c57a-c328-c8b4-d681-ce7298e86ec9.mp3" length="27141595" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《一夫一妻制》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1689</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《送人怀表》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《送人怀表》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e13t4ga</link>
			<guid isPermaLink="false">01acc58b-3705-4c75-bf95-ca77b3e2dff2</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 04 Jul 2021 07:40:10 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/36655050/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2F0356c373-c21b-5e1c-be32-e6cee0ccd016.mp3" length="24445185" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《送人怀表》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1520</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《冷面杀手》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《冷面杀手》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e13t4fp</link>
			<guid isPermaLink="false">3518448c-7532-445d-a0f4-95bceeb5a6d8</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 04 Jul 2021 07:39:47 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/36655033/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-6-4%2Fff46f904-64b6-4b84-f3d3-68f34447213a.mp3" length="33732305" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《冷面杀手》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2100</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《殿下起床》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《殿下起床》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e128jst</link>
			<guid isPermaLink="false">bd039d3f-d87e-487d-967b-eaaee31ec896</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 06 Jun 2021 07:36:13 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/34934109/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-5-6%2F2166aa72-c5c7-b411-1643-d81e7d41b5bf.mp3" length="26351782" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《殿下起床》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1639</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《刷我的卡》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《刷我的卡》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e128jsm</link>
			<guid isPermaLink="false">208cb83d-627b-4375-a9f5-1fbdda401514</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 06 Jun 2021 07:35:51 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/34934102/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-5-6%2Fbd94488f-e9b6-8b10-f74d-80d205707689.mp3" length="32819385" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《刷我的卡》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2043</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《穷人的梦》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《穷人的梦》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e128jsh</link>
			<guid isPermaLink="false">dd3ff730-54d2-42c8-8a35-3e25b0fa4d6a</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 06 Jun 2021 07:35:37 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/34934097/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-5-6%2Fa9662cd2-5b0e-dadd-e8a1-5bbe460b9891.mp3" length="31282622" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《穷人的梦》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1947</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《大艺术家》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《大艺术家》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e128jsd</link>
			<guid isPermaLink="false">3da0fe70-14db-405b-91a8-9063bed79c86</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 06 Jun 2021 07:35:17 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/34934093/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-5-6%2F5989b533-24d7-7c66-628b-c1060adce37e.mp3" length="29858846" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《大艺术家》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1858</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《就是有钱》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《就是有钱》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e128js9</link>
			<guid isPermaLink="false">7063381d-a5e3-4055-b2f7-d148b0158c9c</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 06 Jun 2021 07:35:00 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/34934089/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-5-6%2F94a662d8-3a70-128b-89dc-b712c7f3f78f.mp3" length="21611115" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《就是有钱》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1342</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《为了孩子》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《为了孩子》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e128js3</link>
			<guid isPermaLink="false">8fc26b25-3c7e-4a89-930e-820bafd9f290</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 06 Jun 2021 07:34:42 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/34934083/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-5-6%2F256c10f0-5365-1e98-2b5e-f6037edaf6d9.mp3" length="27139203" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《为了孩子》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1688</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《琴棋书画》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《琴棋书画》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e128jrt</link>
			<guid isPermaLink="false">f0eace59-09da-411b-bc2f-21d0202512f0</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 06 Jun 2021 07:34:22 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/34934077/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-5-6%2Fab5b1a89-6223-cb65-f919-2f84b3153a37.mp3" length="14356896" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《琴棋书画》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>889</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《都不服我》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《都不服我》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e128jrf</link>
			<guid isPermaLink="false">371bab69-5a9a-48bf-961c-4f0d0173000b</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 06 Jun 2021 07:33:49 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/34934063/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-5-6%2Fc7e4854d-1415-e54e-5f79-82a487a1e409.mp3" length="20816359" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《都不服我》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1293</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《穿越神器》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《穿越神器》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e128jr8</link>
			<guid isPermaLink="false">f270f929-4cfe-4cd5-b4bb-8398357da3a0</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 06 Jun 2021 07:33:28 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/34934056/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-5-6%2Fff897b6a-65c8-b176-0386-54760a84e92a.mp3" length="22949156" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《穿越神器》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1426</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无唱助眠】《民以食为天》 高清音频 没有唱]]></title>
			<description><![CDATA[<p>【无唱助眠】郭德纲于谦相声《民以食为天》 高清音频 没有唱</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e128jqu</link>
			<guid isPermaLink="false">bd08b3e5-456a-4d86-a862-ad93cdb5e44a</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 06 Jun 2021 07:33:01 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/34934046/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-5-6%2Fb46ecb60-0ecf-17df-8262-1f4e0716b7ae.mp3" length="34008305" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无唱助眠】郭德纲于谦相声《民以食为天》 高清音频 没有唱&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2117</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo400/14839344/14839344-1619684802588-7857f96383b9.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《你不要》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《你不要》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bjr</link>
			<guid isPermaLink="false">e6665ecf-12a5-4da8-9ce9-d4b625eab606</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:26:33 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861243/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F39dede40-fc9b-a194-93b9-f5325dfc8269.mp3" length="61214208" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《你不要》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3821</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《郭德纲给你咬个怀表》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《郭德纲给你咬个怀表》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bjc</link>
			<guid isPermaLink="false">37b812df-6d0c-4255-8b30-15812d5fe5b5</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:25:51 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861228/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F0d0b536d-ce9d-6a81-fdbd-cc36eb6afc7e.mp3" length="52232948" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《郭德纲给你咬个怀表》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3263</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《不上厕所》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《不上厕所》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bj3</link>
			<guid isPermaLink="false">94817b64-9f90-459c-8f46-d064d636bc02</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:25:27 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861219/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2Fc1332fb6-fe04-c1c9-3dd1-57af0114b2b7.mp3" length="56152516" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《不上厕所》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3496</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《不行》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《不行》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109big</link>
			<guid isPermaLink="false">1a8f5ec7-0e07-4809-af22-7495d6d842ab</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:24:55 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861200/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F0ff0994a-42cf-c3b1-5b08-9638159d4ea8.mp3" length="75300284" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《不行》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4702</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《酸梅汤》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《酸梅汤》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bi6</link>
			<guid isPermaLink="false">1040a7bd-bb8d-4680-a0a1-f60799fb56e4</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:24:22 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861190/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F785eafc9-f5c2-bff7-1f31-38b2c086857a.mp3" length="49008459" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《酸梅汤》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3053</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《五行金木水》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《五行金木水》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bhk</link>
			<guid isPermaLink="false">8ff9dea4-e15a-40bb-ac89-9b32b0c64045</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:23:48 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861172/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2Fc97f58ff-0d2e-bed1-e742-12355affbd63.mp3" length="36729974" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《五行金木水》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2294</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《万圣节》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《万圣节》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bgs</link>
			<guid isPermaLink="false">9c93f048-86cf-4576-b06e-128e284271bc</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:23:08 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861148/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F22c76f81-797e-1bfa-dbea-2daaf2b042ff.mp3" length="38981334" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《万圣节》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2423</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《乾隆谕旨》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《乾隆谕旨》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bgb</link>
			<guid isPermaLink="false">79f0af7e-e7db-49d1-bd13-0c274a46a51e</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:22:18 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861131/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F35e40b06-1d37-fdde-7e68-04ed1bc740a0.mp3" length="27493302" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《乾隆谕旨》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1704</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《夸住宅》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《夸住宅》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bfq</link>
			<guid isPermaLink="false">e4e8a7fb-796f-4e4c-9eab-6af803e317f1</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:21:44 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861114/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2Feb93ca1e-be17-3fb0-e53c-bd0ae2ec6ff2.mp3" length="28224510" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《夸住宅》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1760</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《驴鞭老师》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《驴鞭老师》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bff</link>
			<guid isPermaLink="false">349b5e8e-0ec6-4ee4-9bef-2ac40735b00f</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:21:20 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861103/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2Fbff3c124-63ed-14c9-27d1-5bce5629dd99.mp3" length="49610709" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《驴鞭老师》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3100</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《我要反三俗》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《我要反三俗》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bf7</link>
			<guid isPermaLink="false">1513d466-4494-4701-8a5f-3c8eae566161</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:20:53 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861095/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F70a54542-9500-a65e-bb9b-b0e36ff95bd7.mp3" length="20117416" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《我要反三俗》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1243</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《老老年》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《老老年》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109be8</link>
			<guid isPermaLink="false">3c818e0d-93c7-4923-b9d4-ace60ad4dc8d</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:19:39 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861064/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F5ee26a59-4038-9c99-1b45-feea96fe5fa1.mp3" length="43893486" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《老老年》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2729</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《婚姻与家庭》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《婚姻与家庭》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bdp</link>
			<guid isPermaLink="false">f2df25d6-54d5-46af-900f-6f32f99994bf</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:19:03 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861049/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F14d2d398-f213-2423-7a48-b25ffd1f0f4c.mp3" length="55112473" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《婚姻与家庭》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3430</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《你不要这个样子》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《你不要这个样子》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bd9</link>
			<guid isPermaLink="false">2cc9deab-b26d-4319-b924-8670195dd6d3</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:18:30 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861033/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F12aeca80-c6bf-a5de-3a82-75b7d2b9e4de.mp3" length="61199396" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《你不要这个样子》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3821</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《你要上西天啊》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《你要上西天啊》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bcu</link>
			<guid isPermaLink="false">0c996e3e-4cea-468b-9383-4a8ffda33f8f</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:18:08 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861022/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2Ffaad6931-36f4-d76c-ac4c-6c98415a3099.mp3" length="53459262" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《你要上西天啊》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3340</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【非无损音质】《我要去美国》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【非无损音质】《我要去美国》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bci</link>
			<guid isPermaLink="false">55759fad-049c-491a-9a51-3a921a3de7a5</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:17:44 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32861010/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F7fbcd0d6-2f0f-7541-4214-de3d05f4e4d5.mp3" length="33163526" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【非无损音质】《我要去美国》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2069</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【非无损音质】《讨论个和桑》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【非无损音质】《讨论个和桑》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bc1</link>
			<guid isPermaLink="false">23647366-f5a6-4f79-afc6-bb3c235b4b2d</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:17:22 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32860993/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F8550e273-7750-34ed-aed3-6e8a7e4f8a8e.mp3" length="32695099" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【非无损音质】《讨论个和桑》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2040</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《我要折腾》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《我要折腾》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bbi</link>
			<guid isPermaLink="false">a4801368-703d-41c4-b246-36c3a8e188d8</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:16:55 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32860978/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2Fbaf1b7dd-2209-de22-df28-692962419c5a.mp3" length="57443506" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《我要折腾》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3576</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《爱听相声》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《爱听相声》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bar</link>
			<guid isPermaLink="false">e14ec86d-1283-4e5e-b104-513583fca919</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:16:01 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32860955/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2Ffe9b2000-300a-7099-6fcb-875338425f14.mp3" length="43920407" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《爱听相声》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2741</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《我要旅游》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《我要旅游》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e109bac</link>
			<guid isPermaLink="false">50e6a048-8c12-4a08-803e-097ec85e21fb</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Wed, 05 May 2021 08:15:31 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32860940/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-5%2F9136ebb5-17e0-b37e-947c-f72f7943d20e.mp3" length="28030676" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《我要旅游》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1738</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《卖五器》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《卖五器》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1038li</link>
			<guid isPermaLink="false">da94b72f-bd74-4f97-befc-cea1727f172a</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 02 May 2021 11:14:45 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32661618/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-2%2F130767f9-c8df-fcb5-4b99-bfe3c0034091.mp3" length="49284700" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《卖五器》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3066</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《完美人生》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《完美人生》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1038l6</link>
			<guid isPermaLink="false">f5795bda-ae6f-4ce9-aaa0-25a1b418a3a9</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 02 May 2021 11:14:17 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32661606/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-2%2F299db293-e4ff-fd27-cf78-c60854886b89.mp3" length="25248363" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《完美人生》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1564</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《武松有话说》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《武松有话说》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1038ks</link>
			<guid isPermaLink="false">bf27d94e-b60c-4d88-a983-58685b063ba8</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 02 May 2021 11:13:51 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32661596/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-2%2F6831bcdb-95eb-53b0-29a5-7e445b3378d0.mp3" length="44374636" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《武松有话说》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2759</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《于谦是谁呀》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《于谦是谁呀》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1038ki</link>
			<guid isPermaLink="false">d5cfa1b9-260f-4859-933a-b516fa1ce347</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 02 May 2021 11:13:21 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32661586/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-2%2Ffc9049b6-94d6-ae92-ede0-ce4e0f19fcd3.mp3" length="62729021" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《于谦是谁呀》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3907</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《好朋友》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《好朋友》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1038k8</link>
			<guid isPermaLink="false">867e9a22-9fee-4519-a698-7052fea2941f</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 02 May 2021 11:12:54 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32661576/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-2%2F2b437263-a2ae-73d3-4374-714030faa319.mp3" length="53658020" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《好朋友》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3340</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《我要泡温泉》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《我要泡温泉》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1038jq</link>
			<guid isPermaLink="false">90fc9d4e-b6c4-40c5-b51c-558f69332b75</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 02 May 2021 11:12:23 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32661562/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-2%2F737ea9d0-73b8-215d-470a-b18403471318.mp3" length="26891832" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《我要泡温泉》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1666</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《京中第一名妓》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《京中第一名妓》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1038je</link>
			<guid isPermaLink="false">364e664e-010b-4273-9a4b-53ed4db3c469</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 02 May 2021 11:11:51 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32661550/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-2%2F307b80b8-cb0a-c7ac-f8c0-717c2db8f6d8.mp3" length="26503891" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《京中第一名妓》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1642</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《我是厨神》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《我是厨神》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1038j0</link>
			<guid isPermaLink="false">dbff93dd-d8eb-4d89-9139-a9c5a6affe53</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 02 May 2021 11:11:15 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32661536/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-2%2Fa253042e-dd08-a7ef-5016-fa784f0a1bca.mp3" length="23354422" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《我是厨神》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1446</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《西征梦》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【高清音质】《西征梦》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-e1038i1</link>
			<guid isPermaLink="false">a89c059b-4a9c-493f-9149-8ca9b32ad49b</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Sun, 02 May 2021 11:10:38 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32661505/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-4-2%2F528c41b1-2598-b70e-50e3-3d3469074a42.mp3" length="26371903" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《西征梦》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1634</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《我爱洗澡》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《我爱洗澡》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgjl</link>
			<guid isPermaLink="false">2a7e2268-55e4-403b-8d3c-59a3e1c6be51</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:41:33 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538677/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Fde56b29c-a8f6-2f53-9004-7583d6c7e019.mp3" length="28390898" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《我爱洗澡》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1771</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《总说自己不行》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《总说自己不行》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgj9</link>
			<guid isPermaLink="false">b02ea655-b79e-4c9b-bf49-1584ca67819a</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:41:13 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538665/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Ffae04142-6766-b02b-f481-8b6c6d2cb8df.mp3" length="76283842" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《总说自己不行》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4764</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《于谦的爱人》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《于谦的爱人》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgj1</link>
			<guid isPermaLink="false">ef32aa30-5f50-4fb9-9f3f-63864b71df87</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:40:54 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538657/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F29de0445-6fcc-a844-7f2c-54a5a3aadc3d.mp3" length="26194962" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《于谦的爱人》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1634</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《我和唐僧都一样》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《我和唐僧都一样》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgip</link>
			<guid isPermaLink="false">590a91ae-b1e8-4ed6-a39d-7f15d20b28b3</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:40:32 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538649/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Fe051cd9f-394c-2e5b-8669-11c20bb1480d.mp3" length="32812338" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《我和唐僧都一样》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2037</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《太了不起了》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《太了不起了》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgic</link>
			<guid isPermaLink="false">7aac9992-5414-4eb2-8df6-91e32be7cfc6</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:39:55 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538636/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F94f6587a-09d4-53d3-bec1-25cbd607c3a0.mp3" length="35708331" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《太了不起了》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2228</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《头牌于谦》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《头牌于谦》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgi3</link>
			<guid isPermaLink="false">f8f5b2eb-6762-43b0-ad10-a4198cdfda03</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:39:33 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538627/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F441f71ff-d4dc-48ae-0c5c-fc3c264a7b43.mp3" length="30117915" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《头牌于谦》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1879</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《我去上大学》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《我去上大学》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvghm</link>
			<guid isPermaLink="false">375e461c-ea0d-4335-abc2-76253efab89e</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:39:09 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538614/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Fca032348-9d63-438a-4706-bd3b69cd32f3.mp3" length="45301940" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《我去上大学》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2830</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《老老年》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《老老年》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvghc</link>
			<guid isPermaLink="false">feb59f6f-9e94-4eb7-bec4-0e6ffe1b3a9e</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:38:50 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538604/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Fc9e587f6-1815-a616-cdc4-2e0e33fdc9ed.mp3" length="43592869" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《老老年》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2721</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《走二环》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《走二环》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgh8</link>
			<guid isPermaLink="false">6a0abcb5-9d16-4162-b4c6-a6812e4c30c2</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:38:31 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538600/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Fcb54dbb1-811f-554c-d354-c0f2cabf465e.mp3" length="20516064" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《走二环》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1279</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《重回清朝》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《重回清朝》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvge3</link>
			<guid isPermaLink="false">91347a54-c3a1-4019-98ca-f266633a44c2</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:34:57 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538499/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F562a58c8-9d1b-1fb8-0566-07bb96181245.mp3" length="33488049" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《重回清朝》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2079</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《一亿八千万》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《一亿八千万》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgdn</link>
			<guid isPermaLink="false">957136fd-3959-41c7-a283-a59267c928b5</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:34:26 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538487/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F661ff043-7ae2-0d7c-04d1-3063d708e69b.mp3" length="76240143" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《一亿八千万》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4764</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《你想吃点啥》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《你想吃点啥》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgde</link>
			<guid isPermaLink="false">a3562b3b-fb8c-4568-8561-af9b4f12c4b5</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:34:05 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538478/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Fa758555c-12c1-4c80-bf5f-57502fa352fb.mp3" length="32506488" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《你想吃点啥》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2028</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《东营女子》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《东营女子》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgd8</link>
			<guid isPermaLink="false">0c505ba1-65e6-4e87-bd9f-fa17ae246a4f</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:33:40 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538472/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F0ce0f1c4-ae31-65cd-5882-74b66b9272ed.mp3" length="62417348" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《东营女子》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3897</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《做好事不留名》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《做好事不留名》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgd3</link>
			<guid isPermaLink="false">e93b3081-de12-4b78-8e95-a93af0f7e2fd</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:33:17 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538467/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F2e386ad1-b3e1-ecb5-4eb6-7ff35745f27f.mp3" length="67316005" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《做好事不留名》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4193</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《您这个行业》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《您这个行业》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgcf</link>
			<guid isPermaLink="false">e5b8cd2b-f078-42bd-a27a-de0812bb45e4</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:32:47 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538447/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F45f5f98e-942a-6074-82ad-87c417622f04.mp3" length="54088585" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《您这个行业》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3367</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《脱掉你的衣服》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《脱掉你的衣服》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgc3</link>
			<guid isPermaLink="false">77a36b7b-4455-4734-9564-805f7c21aa4e</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:32:06 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538435/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F4f28bf31-856a-d09d-0b9e-9199a6b3aa1c.mp3" length="71320023" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《脱掉你的衣服》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4444</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《伯父你大爷》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《伯父你大爷》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgb4</link>
			<guid isPermaLink="false">dadf94a3-7020-4d86-96fb-969d958df775</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:30:39 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538404/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F4ada968e-1ffe-fd16-341b-48f18b1de8d0.mp3" length="49101870" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《伯父你大爷》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3055</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《日语专家》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《日语专家》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvgah</link>
			<guid isPermaLink="false">ef50a6e3-0edb-433a-b82c-98912358ce8e</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:30:04 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538385/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F2992299a-cf8c-2c7c-6cb6-e77ba8db89a7.mp3" length="29662533" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《日语专家》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1850</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《舌尖上的大爷》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《舌尖上的大爷》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvga2</link>
			<guid isPermaLink="false">e1e26c58-4d6c-4c42-8d0e-caf009d30341</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:29:31 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538370/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Fbdd342d5-dd6e-912b-a6ab-11453636819c.mp3" length="58019677" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《舌尖上的大爷》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3612</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《李莲英》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《李莲英》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvg9j</link>
			<guid isPermaLink="false">ed3ed69e-ab6b-41fa-baf5-6889668da42d</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:28:55 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538355/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Fb486a641-a65a-791a-d495-4033ee1a9ed7.mp3" length="45125330" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《李莲英》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2807</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《我是厨神》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《我是厨神》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvg91</link>
			<guid isPermaLink="false">4d2d4d99-40a4-4901-b054-9e5dcd0d47ab</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:28:22 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538337/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F24325cab-2463-549d-3818-6dc975aaa1a8.mp3" length="23354422" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《我是厨神》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1446</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《夜盗金陵》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《夜盗金陵》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvg8j</link>
			<guid isPermaLink="false">cb9055cb-cdf1-45f3-9020-70b4aa2591ca</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:27:54 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538323/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F8db3d984-b763-5434-07e1-2166b6f5d88b.mp3" length="35010775" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《夜盗金陵》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2174</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《有翻译》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《有翻译》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvg84</link>
			<guid isPermaLink="false">1976adc9-27fd-4f90-b8be-4179112cfebe</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:27:22 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538308/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Ff121c1a5-5c82-4922-c377-277d6d43b5f2.mp3" length="65774996" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《有翻译》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4097</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《王叔叔》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《王叔叔》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvg7n</link>
			<guid isPermaLink="false">077c9626-6241-4ecf-ad03-60845ff21b6f</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:26:53 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538295/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F8f2d91a5-b374-3a19-5669-15f640563bd1.mp3" length="36921469" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《王叔叔》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2294</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《古典文学》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《古典文学》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvg7e</link>
			<guid isPermaLink="false">6e8bdcc4-e22e-4589-a0f5-9e42befe85a2</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:26:23 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538286/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Fe11538d8-3a97-f444-0ec2-00653a97428b.mp3" length="20698134" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《古典文学》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1279</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《大爷和女鸡》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《大爷和女鸡》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvg73</link>
			<guid isPermaLink="false">6dbcf4ad-cf83-4a27-a759-69adc41d0509</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:25:49 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538275/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F68bd8a37-8ce9-9fc9-e38e-5229e0634866.mp3" length="46934502" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《大爷和女鸡》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2919</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《梦中婚》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《梦中婚》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvg68</link>
			<guid isPermaLink="false">0beb917b-93cd-48d0-97e2-2923cf1dd3bb</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:24:53 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538248/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F04527bf7-0b85-1c83-a501-4c096dfa081b.mp3" length="39160419" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《梦中婚》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2434</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《大爷传奇》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《大爷传奇》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvg5p</link>
			<guid isPermaLink="false">0ca8ccce-9598-4917-9dbb-307505c9a7a6</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:24:16 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538233/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2Fb0f6c6d0-6b5d-0d62-1211-2b983b7fe1ff.mp3" length="55209277" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《大爷传奇》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3437</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《大善人》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《大善人》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvg5i</link>
			<guid isPermaLink="false">6ebb4e7b-47c0-4bdc-95e7-2272abbc6c30</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:23:35 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538226/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F86a2b057-2f4a-aaf5-10e8-12c9a977e036.mp3" length="48230939" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《大善人》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3010</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《门当户对》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《门当户对》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evvg4b</link>
			<guid isPermaLink="false">701ba58d-4377-41a0-b941-f9d6781e11a1</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Fri, 30 Apr 2021 07:22:36 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32538187/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-30%2F38c47951-95a9-e5fc-f14f-d30fa1c310de.mp3" length="33182897" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《门当户对》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2060</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《澡堂子》无唱段清晰无损 郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《澡堂子》无唱段清晰无损 郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtclv</link>
			<guid isPermaLink="false">2b9830f4-3a8c-451c-9c86-881281346936</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 11:13:42 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32469119/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2Fb1af22f9-2cfa-57d7-3474-507a8129aa68.mp3" length="60512497" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《澡堂子》无唱段清晰无损 郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3778</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《说游泳》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《说游泳》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtcll</link>
			<guid isPermaLink="false">4d62f088-b34c-4a36-84a9-94d6d4fe70e2</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 11:13:27 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32469109/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2Fc2562a98-e187-2fc9-d679-15adfc3b6a0a.mp3" length="49433331" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《说游泳》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3085</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《你爸是食神》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《你爸是食神》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtcl3</link>
			<guid isPermaLink="false">a332fa9a-6302-4d6c-9de6-94a2ddd26791</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 11:13:05 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32469091/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F2c73b395-52bc-1808-358b-57143005f29a.mp3" length="36271528" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《你爸是食神》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2253</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《饿了》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《饿了》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtcko</link>
			<guid isPermaLink="false">6c198e53-9781-439e-aa7d-8c672f7ed0f1</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 11:12:48 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32469080/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F0575aff6-3eaa-c8be-f83c-6d48ee98844b.mp3" length="53187383" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《饿了》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3320</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《澡堂传说》郭德纲于谦相声]]></title>
			<description><![CDATA[<p>【高清音质】《澡堂传说》郭德纲于谦相声</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtckg</link>
			<guid isPermaLink="false">b00ba728-fc1f-4661-ac87-fbde29464b9e</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 11:12:31 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32469072/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F2c535966-c855-c43b-4e86-7b978b5eb22b.mp3" length="35699703" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《澡堂传说》郭德纲于谦相声&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2225</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《珍珠鱼》无唱段清晰无损 郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《珍珠鱼》无唱段清晰无损 郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtck3</link>
			<guid isPermaLink="false">e5f20f51-8bf2-4e26-ab0c-6fc6422f2d89</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 11:12:14 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32469059/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F0a7730ed-36d9-6cdf-bde7-433bedc6a134.mp3" length="32139114" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《珍珠鱼》无唱段清晰无损 郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1995</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《鸭子也讲究》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《鸭子也讲究》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtcju</link>
			<guid isPermaLink="false">5e44d50f-6e95-474f-bdf3-68c00b54179b</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 11:11:57 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32469054/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F59b42c98-b253-93d6-69d4-3218194862c0.mp3" length="36852723" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《鸭子也讲究》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2289</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《我这一辈子》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《我这一辈子》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtcjb</link>
			<guid isPermaLink="false">b529d6ef-3022-42bd-8652-5275ba0b7b74</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 11:11:29 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32469035/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F0379c72f-2416-ff8a-c23b-2b463e180dce.mp3" length="15953619" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《我这一辈子》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>993</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《圣诞》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《圣诞》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtcj4</link>
			<guid isPermaLink="false">aead187d-ffa4-4de2-bd5b-4b087dd27e7e</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 11:11:12 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32469028/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2Fec82ce30-2571-91d4-eacc-7552c0819697.mp3" length="61832625" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《圣诞》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3860</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《巴塞罗那的劫匪》无唱段清晰无损郭德綱于謙相聲]]></title>
			<description><![CDATA[<p>【高清音质】《巴塞罗那的劫匪》无唱段清晰无损郭德綱于謙相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtcis</link>
			<guid isPermaLink="false">3fde91a9-c7aa-41d5-bb15-5d9f47b7f7b4</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 11:10:54 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32469020/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F87dc1d03-6641-bbc6-2343-9ee3e805ff5a.mp3" length="32983849" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《巴塞罗那的劫匪》无唱段清晰无损郭德綱于謙相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2048</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《不想去学校》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《不想去学校》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtc1q</link>
			<guid isPermaLink="false">2764b24b-fa0c-4167-9fc5-73abf4517dad</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:53:00 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32468474/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F879d98fd-d19d-23f1-0f40-a0fde5e36345.mp3" length="42706442" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《不想去学校》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2668</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《第一嫖客》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《第一嫖客》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtc1h</link>
			<guid isPermaLink="false">e16d4e7f-f9ad-4c9a-b876-f1c1320c29f3</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:52:43 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32468465/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2Fd5445974-244a-6f1b-9cfe-e7e7043d0adb.mp3" length="28172356" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《第一嫖客》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1757</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【高清音质】《太刺激了 》郭德綱于謙 相聲]]></title>
			<description><![CDATA[<p>【高清音质】《太刺激了 》郭德綱于謙 相聲</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtc16</link>
			<guid isPermaLink="false">a0aed29b-454c-45ea-a62d-c48839932137</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:52:25 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32468454/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F93da4832-abdf-37d9-bd6a-341303123c38.mp3" length="49741177" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【高清音质】《太刺激了 》郭德綱于謙 相聲&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3095</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《上了春晚你变了》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《上了春晚你变了》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtc0s</link>
			<guid isPermaLink="false">d55064b3-0969-4225-958f-dd493314b0ad</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:52:05 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32468444/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F69f6facd-9217-0e9d-9e27-1038ba8c9c46.mp3" length="60119577" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《上了春晚你变了》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3753</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《一滴也没有了》德云社郭德纲于谦相声高清无噪音频版合集·助眠]]></title>
			<description><![CDATA[<p>【无损音质】《一滴也没有了》德云社郭德纲于谦相声高清无噪音频版合集·助眠</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtc0c</link>
			<guid isPermaLink="false">5a0ae5f8-4158-492a-ade2-6d7d675faf31</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:51:48 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32468428/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2Fd9b87edd-ff38-34c9-de25-1e9f726807ec.mp3" length="43955058" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《一滴也没有了》德云社郭德纲于谦相声高清无噪音频版合集·助眠&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2744</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《大爷马桶抠手机》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《大爷马桶抠手机》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtc07</link>
			<guid isPermaLink="false">6350d304-9c4c-4dec-98dc-c6171a4cbc23</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:51:30 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32468423/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F93bd39c5-fe9c-0067-d921-8e23d61aefc3.mp3" length="67249396" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《大爷马桶抠手机》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>4202</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《目瞪口呆》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《目瞪口呆》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtc01</link>
			<guid isPermaLink="false">f114321c-7fac-48e8-b06c-78ab8d14a9f3</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:51:13 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32468417/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F7802293e-8187-8caa-cac0-52ded2b13c9c.mp3" length="48654213" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《目瞪口呆》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3036</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《地方口音》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《地方口音》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtbvn</link>
			<guid isPermaLink="false">ca53e75b-a117-4cd2-a015-0c76bb01b095</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:50:55 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32468407/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F63825f10-4544-cf46-b5d6-7596c34a573d.mp3" length="23891412" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《地方口音》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1489</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《举个栗子》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《举个栗子》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtbul</link>
			<guid isPermaLink="false">691b2e2b-f476-4862-8a37-d5dcd4d00bc8</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:50:28 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32468373/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F318611a5-4c68-203d-4cbd-5acf699c025c.mp3" length="44928303" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《举个栗子》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2806</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《咱不打架》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《咱不打架》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtbl2</link>
			<guid isPermaLink="false">8f6edaf0-a6ab-49e8-acf1-08adffcfb7d7</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:50:05 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32468066/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F0dbc242a-e6b3-2160-98b9-33c543a813ad.mp3" length="51225807" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《咱不打架》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3198</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《奶现交情》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《奶现交情》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtbbc</link>
			<guid isPermaLink="false">42adab4c-7b8a-40c2-9071-c740960274b1</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:49:47 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32467756/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F110e8f14-159e-3b62-823a-e3c1c68d0261.mp3" length="55889009" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《奶现交情》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3488</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《就喜欢动画片》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《就喜欢动画片》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtaun</link>
			<guid isPermaLink="false">27aa6ad7-3348-4321-a43a-de3573fa581f</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:49:21 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32467351/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F6570ae66-c29f-14a9-b7d9-c9a88aca451b.mp3" length="49118883" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《就喜欢动画片》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>3065</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《什么叫贵族》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《什么叫贵族》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evtaiq</link>
			<guid isPermaLink="false">898f7813-9656-44fe-8b09-42e03b452f54</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 10:48:58 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32466970/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F0ba09834-7f13-3338-cf07-d9555bb2d5ca.mp3" length="24915837" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《什么叫贵族》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>1553</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【真·无损音质】《中彩票了》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【真·无损音质】《中彩票了》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evt5rv</link>
			<guid isPermaLink="false">af611477-a078-4354-8ab0-b64b57a2a6aa</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 08:52:19 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32462143/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2Fd1da7923-0767-6031-e68d-c37939881dc0.mp3" length="32778343" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【真·无损音质】《中彩票了》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2044</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《我想要嘘嘘》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《我想要嘘嘘》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evt5lm</link>
			<guid isPermaLink="false">13173bea-78f8-4ec2-8603-e0d52a9d3bf3</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 08:44:36 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32461942/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F9a15fbdc-7162-4df8-5a55-25b834f2674d.mp3" length="43701782" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《我想要嘘嘘》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2728</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
		<item>
			<title><![CDATA[【无损音质】《胡说八道》德云社郭德纲于谦相声高清无噪音频版合集]]></title>
			<description><![CDATA[<p>【无损音质】《胡说八道》德云社郭德纲于谦相声高清无噪音频版合集</p>
]]></description>
			<link>https://anchor.fm/guodegangyuqian/episodes/ep-evt5jh</link>
			<guid isPermaLink="false">dc669c01-df42-4d7c-9684-8f4a513a8e6f</guid>
			<dc:creator><![CDATA[D.M.]]></dc:creator>
			<pubDate>Thu, 29 Apr 2021 08:44:02 GMT</pubDate>
			<enclosure url="https://anchor.fm/s/590ba140/podcast/play/32461873/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-3-29%2F34fb108f-0d57-2fa5-7d95-448c0c67d474.mp3" length="36360641" type="audio/mpeg"/>
			<itunes:summary>&lt;p&gt;【无损音质】《胡说八道》德云社郭德纲于谦相声高清无噪音频版合集&lt;/p&gt;
</itunes:summary>
			<itunes:explicit>No</itunes:explicit>
			<itunes:duration>2269</itunes:duration>
			<itunes:image href="https://d3t3ozftmdmh3i.cloudfront.net/production/podcast_uploaded_nologo/14839344/14839344-1619684806226-e43d7d1df954.jpg"/>
			<itunes:episodeType>full</itunes:episodeType>
		</item>
	</channel>
</rss>
`

func Test_xxx(t *testing.T) {
	flysnowRegexp := regexp.MustCompile(`https://anchor.fm[^\s]*.mp3`)

	params := flysnowRegexp.FindAllStringSubmatch(str, -1)
	fmt.Println(len(params))
	for _, param := range params {
		// fmt.Println(param[0])
		// fmt.Println(regexp.MustCompile(`/https[^\s]*.mp3`).FindString(param[0]))
		mp3 := regexp.MustCompile(`/https[^\s]*.mp3`).FindString(param[0])
		enEscapeUrl, _ := url.QueryUnescape(mp3[1:])
		fmt.Println(enEscapeUrl)
	}
}
