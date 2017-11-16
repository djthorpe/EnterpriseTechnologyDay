
# Enterprise Technology Day

Objectives

  * Create a project and obtain credentials for running API operations 
  * Examine reference material for the Video Intelligence and YouTube Data API's
  * Download SDK's and client libraries
  * Develop a system to load currently running live streams and determine the number of people watching each live stream
  * Develop a system to analyse a video clip and determine knowledge graph classifications of segments within the clip

Code and samples for Enterprise Technology Day. Links:

   * [YouTube Data API](https://developers.google.com/youtube/v3/)
   * [Partner API](https://developers.google.com/youtube/partner/)
   * [Cloud Intelligence API](https://cloud.google.com/video-intelligence/)
   * [Cloud Console](https://console.developers.google.com/)
   * [YouTube Data API Client Libraries](https://developers.google.com/youtube/v3/libraries)
   * [Video Intelligence Client Libraries](https://cloud.google.com/video-intelligence/docs/reference/libraries)
   * [API Explorer](https://developers.google.com/apis-explorer/#p/youtube/v3/)

Here are a couple of publically-acessible videos in Google Cloud Buckets:

  * gs://cloud-ml-sandbox/video/chicago.mp4
  * gs://demomaker/google_gmail.mp4

This is the output I get from the live streaming example:

```
bash$ go run livestreams.go 
+-------------+------------------------------------------------+--------------------+
|     ID      |                     TITLE                      | CONCURRENT VIEWERS |
+-------------+------------------------------------------------+--------------------+
| AQBh9soLSkI | Lofi Hip Hop Radio 24/7 🎧                     |               4074 |
|             | Chill Gaming / Study Beats                     |                    |
| RtU_mdL2vBM | NASA Live - Earth From                         |                382 |
|             | Space (HDVR) ♥ ISS LIVE FEED                   |                    |
|             | #AstronomyDay2017 | Subscribe                  |                    |
|             | now!                                           |                    |
| wwMDvPCGeE0 | NASA TV Public-Education                       |                 77 |
| ezvTXN6vXRM | 24/7 Lofi Hiphop Radio - Beats                 |               1298 |
|             | to Study/Game/Relax                            |                    |
| Fu2etwHzcvw | JapaNews24　～　海外へ日本のニュースをLIVE配信 |                166 |
| HYQk-jXy12w | Town of Jackson Wyoming - SeeJH.com            |                 16 |
| bLhvEJTPXPg | Jackson WY Town Square                         |                 26 |
| iXmRcFeTLY0 | Luton's Teton Cabins - Mount Moran - SeeJH.com |                  1 |
| _PY0x8DeE8Y | Flying Saddle Resort - Snake River in Alpine   |                  3 |
|             | Wyoming - SeeJH.com                            |                    |
| YnclpBYHdYw | Grand Targhee Plaza - SeeJH.com                |                  0 |
| P11y8N22Rq0 | NASA TV  Media                                 |                 23 |
| 2ccaHpy5Ewo | Relaxing Jazz & Bossa Nova Music Radio - 24/7  |               2724 |
|             | Chill Out Piano & Guitar Music Live Stream     |                    |
| Qhq4vQdfrFw | Oxford Martin School Webcam - Broad Street,    |                  4 |
|             | Oxford                                         |                    |
| XbOuldSBnu4 | ⚡ FALLOUT RADIO - Playing All Radio Stations  |                122 |
|             | (DJ's Included!) 🔴 24/7                       |                    |
| q8Sg0RL5fBU | Henry's Fork - Island Park, Idaho - SeeJH.com  |                  3 |
| 37gYYw0L6CY | VINTAGE RADIO (music from Fallout, Bioshock,   |                 28 |
|             | Mafia and more!)                               |                    |
| p5mu_febWXg | Brooks Falls - Katmai National Park, Alaska    |                 59 |
|             | powered by EXPLORE.org                         |                    |
| Rn3Nn0Eqw5U | GTV Live                                       |                305 |
| GCP97SY0pMI | River Watch - Katmai National Park, Alaska     |                  4 |
|             | powered by EXPLORE.org                         |                    |
| xIpFO5Faz5U | UKF Drum & Bass 24/7 Mix - Liquid Radio 🎧     |                380 |
+-------------+------------------------------------------------+--------------------+
```

This is the output I get from the video annotation example:

```
bash$ go run vi-analyse.go gs://demomaker/google_gmail.mp4
+---------------+------------+--------------------------------+--------------+--------------+------------+
|     TYPE      |   ENTITY   |          DESCRIPTION           |    START     |     END      | CONFIDENCE |
+---------------+------------+--------------------------------+--------------+--------------+------------+
| shot_label    | /m/07s6nbt | text                           | 0s           | 3.032998s    |  0.4245241 |
| shot_label    | /m/02y3rj  | presentation > person          | 36.132198s   | 39.304276s   | 0.62344676 |
|               |            |                                | 39.339046s   | 41.134706s   |  0.6665199 |
|               |            |                                | 41.169476s   | 43.539842s   |  0.7542678 |
|               |            |                                | 43.574612s   | 45.038956s   | 0.82267106 |
|               |            |                                | 45.073726s   | 46.380604s   |  0.8339506 |
|               |            |                                | 46.415374s   | 46.92054s    |  0.8603124 |
|               |            |                                | 46.938926s   | 47.844948s   | 0.84523356 |
|               |            |                                | 47.879718s   | 51.818738s   | 0.82150495 |
|               |            |                                | 51.853508s   | 52.777916s   |  0.8944215 |
|               |            |                                | 1m13.309498s | 1m17.77207s  | 0.65665555 |
|               |            |                                | 1m26.347512s | 1m30.024756s | 0.68976253 |
| shot_label    | /m/085n4   | website > weak entity          | 43.574612s   | 45.038956s   | 0.50376886 |
|               |            |                                | 46.415374s   | 46.92054s    | 0.51534617 |
|               |            |                                | 46.938926s   | 47.844948s   | 0.51419723 |
|               |            |                                | 47.879718s   | 51.818738s   |  0.4996919 |
|               |            |                                | 51.853508s   | 52.777916s   |  0.5037447 |
| shot_label    | /m/013s93  | t shirt > top                  | 3.067768s    | 7.076328s    | 0.62171304 |
|               |            |                                | 7.111098s    | 15.774466s   | 0.90612245 |
|               |            |                                | 19.678716s   | 27.556756s   |  0.6206226 |
|               |            |                                | 52.812686s   | 58.721216s   | 0.61108226 |
|               |            |                                | 1m38.426348s | 1m45.10182s  |  0.6239333 |
| shot_label    | /m/0dzd8   | neck > human body              | 7.111098s    | 15.774466s   | 0.41365626 |
| shot_label    | /m/0n5szg6 | business executive > person    | 1m4.36797s   | 1m13.274728s |  0.5802008 |
|               |            |                                | 1m17.80684s  | 1m20.839838s |  0.5888127 |
| shot_label    | /m/083jv   | white                          | 1m45.13659s  | 1m46.478238s |  0.4450501 |
| shot_label    | /m/062zr   | public relations > person      | 46.415374s   | 46.92054s    |  0.4708813 |
| shot_label    | /m/03g09t  | clip art > graphics            | 1m46.513008s | 1m52.212918s | 0.40146247 |
| shot_label    | /m/02v0m2  | diagram                        | 33.16874s    | 36.097428s   | 0.46559462 |
| shot_label    | /m/01mf0   | software                       | 39.339046s   | 41.134706s   |  0.4289541 |
| shot_label    | /m/021sdg  | graphics > artwork             | 1m45.13659s  | 1m46.478238s |  0.5937147 |
|               |            |                                | 1m46.513008s | 1m52.212918s |  0.8131164 |
| shot_label    | /m/02cwm   | design                         | 27.591526s   | 33.13397s    |  0.6150597 |
|               |            |                                | 33.16874s    | 36.097428s   |  0.6150597 |
| shot_label    | /m/01vkl   | circle > shape                 | 27.591526s   | 33.13397s    | 0.59355646 |
| shot_label    | /m/03bxgrp | company > organization         | 41.169476s   | 43.539842s   |  0.4619668 |
|               |            |                                | 43.574612s   | 45.038956s   |  0.5327454 |
|               |            |                                | 45.073726s   | 46.380604s   | 0.62085783 |
|               |            |                                | 46.415374s   | 46.92054s    |  0.6537138 |
|               |            |                                | 46.938926s   | 47.844948s   |  0.5506257 |
|               |            |                                | 47.879718s   | 51.818738s   | 0.44909155 |
|               |            |                                | 51.853508s   | 52.777916s   | 0.45067522 |
| shot_label    | /m/0mwc    | angle                          | 33.16874s    | 36.097428s   |  0.4649368 |
| shot_label    | /m/0jyfg   | glasses                        | 1m4.36797s   | 1m13.274728s | 0.56006306 |
|               |            |                                | 1m17.80684s  | 1m20.839838s |  0.5589393 |
| shot_label    | /m/05b1rx  | online advertising >           | 43.574612s   | 45.038956s   | 0.43043193 |
|               |            | advertising                    |              |              |            |
|               |            |                                | 46.415374s   | 46.92054s    |  0.5977116 |
|               |            |                                | 46.938926s   | 47.844948s   |  0.4830443 |
| shot_label    | /m/01qkbx  | professional > person          | 41.169476s   | 43.539842s   |  0.6610328 |
|               |            |                                | 43.574612s   | 45.038956s   | 0.73607576 |
|               |            |                                | 45.073726s   | 46.380604s   |  0.7499733 |
|               |            |                                | 46.415374s   | 46.92054s    |        0.6 |
|               |            |                                | 46.938926s   | 47.844948s   |  0.7334697 |
|               |            |                                | 51.853508s   | 52.777916s   |  0.4465192 |
|               |            |                                | 1m4.36797s   | 1m13.274728s |  0.9135862 |
|               |            |                                | 1m13.309498s | 1m17.77207s  |  0.6472558 |
|               |            |                                | 1m17.80684s  | 1m20.839838s |  0.9272224 |
|               |            |                                | 1m26.347512s | 1m30.024756s | 0.41451982 |
| shot_label    | /m/0643t   | personal computer > computer   | 15.809236s   | 17.848286s   | 0.57247293 |
| shot_label    | /m/0dwx7   | logo > graphics                | 1m45.13659s  | 1m46.478238s |  0.4779007 |
|               |            |                                | 1m46.513008s | 1m52.212918s |  0.5097802 |
| shot_label    | /m/01m2v   | computer keyboard > computer   | 15.809236s   | 17.848286s   | 0.60816985 |
| shot_label    | /m/0p9xx   | visual arts > art              | 27.591526s   | 33.13397s    | 0.45865557 |
| shot_label    | /m/091410  | collar > clothing              | 7.111098s    | 15.774466s   | 0.45955557 |
| shot_label    | /m/03qh03g | media                          | 46.415374s   | 46.92054s    | 0.42600083 |
| shot_label    | /m/01h8n0  | conversation > communication   | 36.132198s   | 39.304276s   |  0.4518466 |
|               |            |                                | 39.339046s   | 41.134706s   | 0.44234124 |
|               |            |                                | 41.169476s   | 43.539842s   | 0.45186457 |
|               |            |                                | 43.574612s   | 45.038956s   | 0.46615928 |
|               |            |                                | 45.073726s   | 46.380604s   | 0.57594925 |
|               |            |                                | 46.415374s   | 46.92054s    |  0.4656734 |
|               |            |                                | 46.938926s   | 47.844948s   | 0.47492284 |
|               |            |                                | 47.879718s   | 51.818738s   | 0.46630585 |
|               |            |                                | 51.853508s   | 52.777916s   | 0.58575684 |
| shot_label    | /m/01gq53  | performance > event,           | 1m20.874608s | 1m26.312742s | 0.55028087 |
|               |            | entertainment                  |              |              |            |
|               |            |                                | 1m30.059526s | 1m38.391578s | 0.62056136 |
| shot_label    | /m/02wzbmj | standing > person              | 7.111098s    | 15.774466s   |  0.5869962 |
|               |            |                                | 19.678716s   | 27.556756s   |  0.5588316 |
|               |            |                                | 52.812686s   | 58.721216s   |  0.5256574 |
|               |            |                                | 1m26.347512s | 1m30.024756s |    0.42976 |
|               |            |                                | 1m38.426348s | 1m45.10182s  | 0.55333024 |
| shot_label    | /m/074ft   | song > music                   | 1m20.874608s | 1m26.312742s |  0.5637414 |
| shot_label    | /m/03c31   | graphic design > artwork       | 1m46.513008s | 1m52.212918s | 0.41454726 |
| shot_label    | /m/03zyjs  | polo shirt > top               | 7.111098s    | 15.774466s   |  0.5870668 |
| shot_label    | /m/09s1f   | business > organization        | 36.132198s   | 39.304276s   |  0.6444387 |
|               |            |                                | 39.339046s   | 41.134706s   |  0.5669425 |
|               |            |                                | 41.169476s   | 43.539842s   | 0.81588036 |
|               |            |                                | 43.574612s   | 45.038956s   |  0.9111362 |
|               |            |                                | 45.073726s   | 46.380604s   |          1 |
|               |            |                                | 46.415374s   | 46.92054s    |          1 |
|               |            |                                | 46.938926s   | 47.844948s   |          1 |
|               |            |                                | 47.879718s   | 51.818738s   |          1 |
|               |            |                                | 51.853508s   | 52.777916s   |  0.8892581 |
|               |            |                                | 58.755986s   | 1m4.3332s    | 0.43920207 |
|               |            |                                | 1m4.36797s   | 1m13.274728s |  0.5669425 |
|               |            |                                | 1m13.309498s | 1m17.77207s  | 0.84441346 |
|               |            |                                | 1m17.80684s  | 1m20.839838s |  0.5669425 |
|               |            |                                | 1m26.347512s | 1m30.024756s | 0.86101973 |
| shot_label    | /m/015lz1  | singing > entertainment,       | 1m20.874608s | 1m26.312742s |  0.5370969 |
|               |            | person                         |              |              |            |
|               |            |                                | 1m30.059526s | 1m38.391578s |  0.6809759 |
| shot_label    | /m/01n4qj  | shirt > top                    | 3.067768s    | 7.076328s    |   0.476455 |
|               |            |                                | 7.111098s    | 15.774466s   | 0.75102776 |
|               |            |                                | 19.678716s   | 27.556756s   |   0.476455 |
|               |            |                                | 1m38.426348s | 1m45.10182s  |   0.476455 |
| shot_label    | /m/01mfj   | computer hardware > computer   | 15.809236s   | 17.848286s   |  0.4062668 |
| shot_label    | /m/086nh   | web page                       | 36.132198s   | 39.304276s   | 0.54652894 |
|               |            |                                | 39.339046s   | 41.134706s   |    0.48516 |
|               |            |                                | 43.574612s   | 45.038956s   |   0.616773 |
|               |            |                                | 45.073726s   | 46.380604s   |    0.48516 |
|               |            |                                | 46.415374s   | 46.92054s    |  0.6369214 |
|               |            |                                | 46.938926s   | 47.844948s   |  0.6385848 |
|               |            |                                | 47.879718s   | 51.818738s   |        0.6 |
|               |            |                                | 51.853508s   | 52.777916s   |    0.48516 |
| shot_label    | /m/062581  | sleeve > clothing              | 7.111098s    | 15.774466s   | 0.47732645 |
| shot_label    | /m/0641k   | paper > material               | 27.591526s   | 33.13397s    |  0.6071224 |
|               |            |                                | 33.16874s    | 36.097428s   | 0.51375073 |
| shot_label    | /m/012t_z  | businessperson > person        | 45.073726s   | 46.380604s   | 0.68911326 |
|               |            |                                | 46.415374s   | 46.92054s    | 0.57391524 |
|               |            |                                | 1m4.36797s   | 1m13.274728s | 0.74188906 |
|               |            |                                | 1m17.80684s  | 1m20.839838s | 0.75472295 |
| shot_label    | /m/02csf   | drawing > artwork              | 27.591526s   | 33.13397s    |  0.9392641 |
|               |            |                                | 33.16874s    | 36.097428s   |  0.8999082 |
| shot_label    | /m/01kq3x  | white collar worker > person   | 45.073726s   | 46.380604s   |  0.4041419 |
|               |            |                                | 1m4.36797s   | 1m13.274728s | 0.48669523 |
|               |            |                                | 1m17.80684s  | 1m20.839838s |  0.4857711 |
| shot_label    | /m/0191f8  | learning > person              | 17.883056s   | 19.66033s    |  0.5059999 |
| shot_label    | /m/07glzq  | sketch > drawing               | 33.16874s    | 36.097428s   | 0.65242344 |
| shot_label    | /m/068k4   | public speaking > person       | 45.073726s   | 46.380604s   | 0.48122337 |
|               |            |                                | 46.415374s   | 46.92054s    |  0.4813425 |
| segment_label | /m/02wzbmj | standing > person              | 0s           | 1m52.212918s |  0.4038386 |
| segment_label | /m/013s93  | t shirt > top                  | 0s           | 1m52.212918s |  0.4240086 |
| segment_label | /m/09s1f   | business > organization        | 0s           | 1m52.212918s | 0.77872217 |
+---------------+------------+--------------------------------+--------------+--------------+------------+
```

