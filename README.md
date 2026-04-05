**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-04-05 00:34:37 UTC（2026-04-05 08:34:37 UTC+8）

**代理总数：375**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 375 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 375 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.34.179.58:8446 | ✓ 660ms | ✓ 937ms | ✓ 759ms | ✓ 897ms | ✓ 723ms | http |
| 208.87.243.199:7878 | ✓ 387ms | ✓ 1042ms | ✓ 1078ms | ✓ 887ms | ✓ 688ms | http |
| 38.145.208.215:8444 | ✓ 684ms | ✓ 930ms | ✓ 780ms | ✓ 910ms | ✓ 720ms | http |
| 38.34.179.29:8452 | ✓ 706ms | ✓ 1125ms | ✓ 525ms | ✓ 942ms | ✓ 735ms | http |
| 38.34.179.26:8446 | ✓ 660ms | ✓ 1220ms | ✓ 477ms | ✓ 937ms | ✓ 719ms | http |
| 45.136.131.28:8449 | ✓ 660ms | ✓ 919ms | ✓ 928ms | ✓ 910ms | ✓ 731ms | http |
| 38.145.218.235:8453 | ✓ 794ms | ✓ 955ms | ✓ 608ms | ✓ 920ms | ✓ 746ms | http |
| 38.145.203.19:8447 | ✓ 684ms | ✓ 850ms | ✓ 974ms | ✓ 1001ms | ✓ 737ms | http |
| 45.136.131.59:8444 | ✓ 792ms | ✓ 998ms | ✓ 565ms | ✓ 1009ms | ✓ 761ms | http |
| 38.145.218.9:8445 | ✓ 684ms | ✓ 925ms | ✓ 898ms | ✓ 1229ms | ✓ 696ms | http |
| 38.145.220.27:8445 | ✓ 660ms | ✓ 1198ms | ✓ 497ms | ✓ 992ms | ✓ 820ms | http |
| 45.136.131.57:8449 | ✓ 660ms | ✓ 1066ms | ✓ 1242ms | ✓ 979ms | ✓ 749ms | http |
| 38.145.220.32:8449 | ✓ 660ms | ✓ 1155ms | ✓ 540ms | ✓ 1065ms | ✓ 829ms | http |
| 35.225.22.61:80 | ✓ 660ms | ✓ 1494ms | ✓ 1044ms | ✓ 1039ms | ✓ 923ms | http |
| 38.34.179.16:8452 | ✓ 684ms | ✓ 981ms | ✓ 1301ms | ✓ 1143ms | ✓ 756ms | http |
| 38.145.203.108:8452 | ✓ 1197ms | ✓ 873ms | ✓ 438ms | ✓ 1104ms | ✓ 1335ms | http |
| 38.145.220.65:8446 | ✓ 659ms | ✓ 949ms | ✓ 747ms | ✓ 1201ms | ✓ 1172ms | http |
| 38.34.178.152:8451 | ✓ 683ms | ✓ 1012ms | ✓ 968ms | ✓ 1016ms | ✓ 882ms | http |
| 45.136.131.51:8446 | ✓ 660ms | ✓ 1579ms | ✓ 1012ms | ✓ 971ms | ✓ 726ms | http |
| 38.145.203.111:8450 | ✓ 1196ms | ✓ 895ms | ✓ 723ms | ✓ 1347ms | ✓ 953ms | http |
| 38.145.220.36:8451 | ✓ 661ms | ✓ 1201ms | ✓ 493ms | ✓ 1016ms | ✓ 1149ms | http |
| 38.145.220.15:8451 | ✓ 659ms | ✓ 1267ms | ✓ 428ms | ✓ 960ms | ✓ 1144ms | http |
| 38.145.218.160:8448 | 否 | ✓ 1048ms | ✓ 293ms | ✓ 1018ms | ✓ 839ms | http |
| 38.145.218.162:8448 | 否 | ✓ 1077ms | ✓ 365ms | ✓ 999ms | ✓ 833ms | http |
| 45.136.130.182:8446 | ✓ 799ms | ✓ 978ms | ✓ 1037ms | ✓ 1645ms | ✓ 726ms | http |
| 38.145.203.106:8448 | ✓ 1555ms | ✓ 1093ms | ✓ 708ms | ✓ 1503ms | ✓ 729ms | http |
| 38.34.179.24:8447 | 否 | ✓ 1952ms | ✓ 383ms | ✓ 964ms | ✓ 764ms | http |
| 38.34.179.87:8447 | ✓ 662ms | ✓ 1461ms | ✓ 693ms | ✓ 959ms | ✓ 1002ms | http |
| 45.136.130.178:8453 | ✓ 793ms | ✓ 1055ms | ✓ 965ms | ✓ 1888ms | ✓ 822ms | http |
| 38.34.179.186:8444 | ✓ 660ms | ✓ 1183ms | ✓ 1401ms | ✓ 969ms | ✓ 1166ms | http |
| 147.161.210.140:8800 | ✓ 929ms | ✓ 1338ms | ✓ 1189ms | 否 | ✓ 1156ms | http |
| 37.187.109.70:10111 | ✓ 1055ms | ✓ 1263ms | ✓ 1227ms | 否 | ✓ 1695ms | http |
| 115.231.181.40:8128 | ✓ 1087ms | ✓ 1281ms | ✓ 1040ms | ✓ 1300ms | ✓ 1156ms | http |
| 95.213.217.168:52004 | ✓ 1085ms | ✓ 1493ms | 否 | ✓ 1505ms | ✓ 1077ms | http |
| 185.191.236.162:3128 | ✓ 1097ms | ✓ 1485ms | ✓ 1530ms | ✓ 1492ms | 否 | http |
| 34.96.238.40:8080 | 否 | ✓ 1461ms | ✓ 1420ms | ✓ 1524ms | ✓ 1508ms | http |
| 1.231.81.166:3128 | ✓ 987ms | ✓ 1178ms | 否 | ✓ 1292ms | ✓ 1389ms | http |
| 38.145.218.216:8449 | ✓ 1196ms | ✓ 1434ms | ✓ 1227ms | ✓ 1478ms | ✓ 1158ms | http |
| 218.108.131.186:17890 | ✓ 975ms | ✓ 1242ms | ✓ 1024ms | ✓ 1321ms | ✓ 1060ms | http |
| 111.227.254.9:22222 | ✓ 1252ms | ✓ 1557ms | ✓ 1268ms | ✓ 1577ms | ✓ 1295ms | http |
| 113.160.132.26:8080 | ✓ 1605ms | ✓ 1570ms | ✓ 1482ms | ✓ 1444ms | ✓ 1114ms | http |
| 111.227.254.12:22222 | ✓ 1286ms | ✓ 1549ms | ✓ 1276ms | ✓ 1622ms | ✓ 1275ms | http |
| 159.223.71.162:443 | ✓ 1620ms | 否 | ✓ 1865ms | ✓ 1279ms | ✓ 1029ms | http |
| 159.223.71.162:8080 | ✓ 1621ms | 否 | ✓ 1877ms | ✓ 1281ms | ✓ 1010ms | http |
| 212.58.132.5:8888 | ✓ 1622ms | 否 | ✓ 1465ms | ✓ 1594ms | ✓ 1205ms | http |
| 167.103.115.102:8800 | ✓ 1619ms | 否 | ✓ 1322ms | ✓ 1331ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1732ms | 否 | 否 | ✓ 1987ms | ✓ 1870ms | http |
| 150.241.71.15:1080 | ✓ 1068ms | ✓ 1849ms | ✓ 897ms | ✓ 1283ms | 否 | http |
| 38.145.208.173:8444 | ✓ 313ms | ✓ 841ms | ✓ 307ms | ✓ 887ms | ✓ 678ms | http |
| 38.145.218.227:8445 | ✓ 298ms | ✓ 889ms | ✓ 283ms | ✓ 858ms | ✓ 718ms | http |
| 38.34.179.50:8444 | ✓ 313ms | ✓ 877ms | ✓ 305ms | ✓ 961ms | ✓ 714ms | http |
| 38.145.208.174:8444 | ✓ 318ms | ✓ 839ms | ✓ 282ms | ✓ 959ms | ✓ 705ms | http |
| 38.145.208.179:8450 | ✓ 359ms | ✓ 877ms | ✓ 394ms | ✓ 909ms | ✓ 681ms | http |
| 38.145.208.175:8447 | ✓ 273ms | ✓ 962ms | ✓ 281ms | ✓ 933ms | ✓ 686ms | http |
| 38.145.208.223:8450 | ✓ 390ms | ✓ 898ms | ✓ 253ms | ✓ 934ms | ✓ 759ms | http |
| 38.145.208.220:8448 | ✓ 363ms | ✓ 890ms | ✓ 259ms | ✓ 942ms | ✓ 767ms | http |
| 38.34.179.53:8451 | ✓ 303ms | ✓ 850ms | ✓ 284ms | ✓ 1060ms | ✓ 697ms | http |
| 38.34.179.49:8445 | ✓ 292ms | ✓ 935ms | ✓ 310ms | ✓ 950ms | ✓ 705ms | http |
| 38.145.208.180:8451 | ✓ 412ms | ✓ 870ms | ✓ 376ms | ✓ 927ms | ✓ 737ms | http |
| 38.145.208.183:8448 | ✓ 349ms | ✓ 1021ms | ✓ 316ms | ✓ 877ms | ✓ 695ms | http |
| 38.145.208.184:8444 | ✓ 420ms | ✓ 988ms | ✓ 290ms | ✓ 900ms | ✓ 717ms | http |
| 38.145.218.14:8446 | ✓ 368ms | ✓ 995ms | ✓ 313ms | ✓ 900ms | ✓ 703ms | http |
| 38.34.179.37:8450 | ✓ 490ms | ✓ 885ms | ✓ 285ms | ✓ 917ms | ✓ 709ms | http |
| 38.34.179.102:8444 | ✓ 457ms | ✓ 858ms | ✓ 288ms | ✓ 1015ms | ✓ 700ms | http |
| 38.34.179.96:8446 | ✓ 436ms | ✓ 864ms | ✓ 307ms | ✓ 990ms | ✓ 718ms | http |
| 38.34.179.61:8445 | ✓ 433ms | ✓ 850ms | ✓ 301ms | ✓ 1011ms | ✓ 711ms | http |
| 38.34.179.97:8446 | ✓ 429ms | ✓ 862ms | ✓ 313ms | ✓ 983ms | ✓ 700ms | http |
| 38.34.179.39:8452 | ✓ 503ms | ✓ 936ms | ✓ 303ms | ✓ 928ms | ✓ 714ms | http |
| 38.34.179.202:8448 | ✓ 396ms | ✓ 1033ms | ✓ 296ms | ✓ 968ms | ✓ 705ms | http |
| 38.34.179.40:8446 | ✓ 534ms | ✓ 860ms | ✓ 312ms | ✓ 968ms | ✓ 737ms | http |
| 38.34.179.194:8451 | ✓ 364ms | ✓ 1012ms | ✓ 303ms | ✓ 935ms | ✓ 702ms | http |
| 38.34.179.95:8444 | ✓ 490ms | ✓ 882ms | ✓ 290ms | ✓ 1014ms | ✓ 730ms | http |
| 38.34.179.228:8453 | ✓ 389ms | ✓ 1021ms | ✓ 291ms | ✓ 957ms | ✓ 723ms | http |
| 38.34.179.101:8453 | ✓ 452ms | ✓ 897ms | ✓ 303ms | ✓ 1034ms | ✓ 712ms | http |
| 38.34.179.103:8448 | ✓ 468ms | ✓ 888ms | ✓ 314ms | ✓ 1016ms | ✓ 728ms | http |
| 38.145.203.135:8444 | ✓ 429ms | ✓ 860ms | ✓ 612ms | ✓ 909ms | ✓ 698ms | http |
| 38.34.179.100:8453 | ✓ 525ms | ✓ 860ms | ✓ 280ms | ✓ 1072ms | ✓ 713ms | http |
| 38.34.179.203:8446 | ✓ 412ms | ✓ 1052ms | ✓ 287ms | ✓ 962ms | ✓ 699ms | http |
| 38.34.179.80:8452 | ✓ 451ms | ✓ 893ms | ✓ 566ms | ✓ 916ms | ✓ 697ms | http |
| 38.34.179.99:8446 | ✓ 504ms | ✓ 1042ms | ✓ 370ms | ✓ 971ms | ✓ 718ms | http |
| 38.34.179.76:8444 | ✓ 438ms | ✓ 858ms | ✓ 623ms | ✓ 938ms | ✓ 757ms | http |
| 38.34.179.82:8447 | ✓ 419ms | ✓ 858ms | ✓ 591ms | ✓ 971ms | ✓ 781ms | http |
| 38.34.179.75:8447 | ✓ 403ms | ✓ 934ms | ✓ 545ms | ✓ 961ms | ✓ 790ms | http |
| 38.34.179.74:8447 | ✓ 429ms | ✓ 874ms | ✓ 590ms | ✓ 978ms | ✓ 792ms | http |
| 38.145.208.239:8448 | ✓ 478ms | ✓ 928ms | ✓ 814ms | ✓ 922ms | ✓ 716ms | http |
| 38.145.208.229:8453 | ✓ 454ms | ✓ 837ms | ✓ 892ms | ✓ 934ms | ✓ 707ms | http |
| 38.145.208.228:8445 | ✓ 541ms | ✓ 838ms | ✓ 870ms | ✓ 935ms | ✓ 728ms | http |
| 38.145.208.226:8448 | ✓ 476ms | ✓ 833ms | ✓ 888ms | ✓ 921ms | ✓ 746ms | http |
| 38.34.179.32:8443 | ✓ 474ms | 否 | ✓ 288ms | ✓ 906ms | ✓ 731ms | http |
| 38.34.179.35:8443 | ✓ 504ms | 否 | ✓ 282ms | ✓ 946ms | ✓ 709ms | http |
| 38.34.179.46:8443 | ✓ 464ms | 否 | ✓ 293ms | ✓ 949ms | ✓ 701ms | http |
| 38.34.179.33:8443 | ✓ 487ms | 否 | ✓ 316ms | ✓ 908ms | ✓ 720ms | http |
| 38.34.179.34:8443 | ✓ 500ms | 否 | ✓ 289ms | ✓ 918ms | ✓ 749ms | http |
| 38.34.179.78:8448 | ✓ 388ms | ✓ 869ms | ✓ 1254ms | ✓ 1319ms | ✓ 690ms | http |
| 38.34.179.105:8449 | ✓ 915ms | ✓ 941ms | ✓ 317ms | ✓ 1152ms | ✓ 749ms | http |
| 38.34.179.89:8449 | ✓ 341ms | ✓ 961ms | ✓ 443ms | ✓ 1270ms | ✓ 737ms | http |
| 38.145.203.35:8450 | ✓ 413ms | ✓ 1005ms | ✓ 671ms | ✓ 1289ms | ✓ 815ms | http |
| 38.34.179.106:8445 | ✓ 318ms | ✓ 1196ms | ✓ 322ms | ✓ 1042ms | ✓ 1135ms | http |
| 38.34.179.63:8448 | ✓ 388ms | ✓ 977ms | ✓ 1041ms | ✓ 927ms | ✓ 1209ms | http |
| 38.145.203.39:8445 | ✓ 1599ms | 否 | ✓ 287ms | ✓ 961ms | ✓ 823ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
