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

最后更新：2026-03-21 20:23:43 UTC（2026-03-22 04:23:43 UTC+8）

**代理总数：185**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 185 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 185 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.34.179.75:8453 | ✓ 378ms | ✓ 683ms | ✓ 399ms | ✓ 1198ms | ✓ 1611ms | http |
| 147.161.210.140:8800 | ✓ 1474ms | ✓ 1121ms | ✓ 865ms | ✓ 1043ms | ✓ 1673ms | http |
| 167.103.34.108:8800 | ✓ 1353ms | 否 | ✓ 1460ms | ✓ 1630ms | 否 | http |
| 38.145.203.19:8449 | ✓ 511ms | ✓ 1839ms | ✓ 1558ms | 否 | 否 | http |
| 38.34.179.54:8446 | 否 | ✓ 1567ms | 否 | ✓ 1657ms | ✓ 1342ms | http |
| 38.34.179.150:8449 | ✓ 203ms | ✓ 664ms | ✓ 247ms | ✓ 864ms | ✓ 596ms | http |
| 38.34.179.162:8451 | ✓ 236ms | ✓ 754ms | ✓ 147ms | ✓ 745ms | ✓ 654ms | http |
| 38.34.179.73:8446 | ✓ 209ms | ✓ 616ms | ✓ 574ms | ✓ 749ms | ✓ 680ms | http |
| 38.34.179.74:8449 | ✓ 265ms | ✓ 636ms | ✓ 505ms | ✓ 744ms | ✓ 1357ms | http |
| 38.34.179.40:8446 | ✓ 700ms | ✓ 1478ms | ✓ 157ms | ✓ 820ms | ✓ 674ms | http |
| 38.34.179.83:8448 | ✓ 280ms | ✓ 829ms | ✓ 727ms | ✓ 1802ms | ✓ 510ms | http |
| 43.99.54.236:5555 | ✓ 638ms | ✓ 924ms | ✓ 684ms | ✓ 826ms | ✓ 644ms | http |
| 45.136.130.177:8448 | ✓ 1065ms | ✓ 940ms | ✓ 639ms | ✓ 1115ms | 否 | http |
| 38.145.208.181:8445 | ✓ 1175ms | ✓ 1696ms | ✓ 164ms | ✓ 804ms | ✓ 1697ms | http |
| 35.225.22.61:80 | ✓ 794ms | 否 | ✓ 895ms | ✓ 1341ms | ✓ 1224ms | http |
| 38.34.179.203:8451 | ✓ 1849ms | ✓ 714ms | ✓ 169ms | ✓ 969ms | 否 | http |
| 45.140.147.155:1081 | ✓ 587ms | ✓ 1436ms | ✓ 931ms | 否 | ✓ 1349ms | http |
| 45.140.147.155:1082 | ✓ 581ms | ✓ 1516ms | ✓ 903ms | 否 | ✓ 1341ms | http |
| 38.34.183.47:8452 | ✓ 281ms | ✓ 1274ms | ✓ 1784ms | ✓ 657ms | ✓ 672ms | http |
| 38.34.178.155:8448 | 否 | ✓ 1923ms | 否 | ✓ 1029ms | ✓ 544ms | http |
| 121.126.185.63:25152 | ✓ 1219ms | 否 | ✓ 1440ms | ✓ 1639ms | 否 | http |
| 45.144.28.81:10808 | ✓ 703ms | ✓ 1529ms | ✓ 754ms | 否 | 否 | http |
| 38.34.178.245:8446 | ✓ 177ms | ✓ 686ms | 否 | ✓ 1579ms | ✓ 908ms | http |
| 38.34.179.8:8449 | ✓ 540ms | ✓ 1087ms | 否 | ✓ 1411ms | ✓ 1032ms | http |
| 38.34.179.19:8449 | ✓ 573ms | ✓ 1100ms | 否 | ✓ 1419ms | ✓ 1064ms | http |
| 106.75.15.167:7890 | ✓ 1806ms | ✓ 1447ms | 否 | ✓ 1863ms | ✓ 1365ms | http |
| 219.117.204.211:7799 | ✓ 742ms | ✓ 1341ms | ✓ 1844ms | 否 | 否 | http |
| 5.129.237.45:49488 | ✓ 794ms | ✓ 1779ms | 否 | 否 | ✓ 1424ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1658ms | ✓ 1107ms | ✓ 1325ms | http |
| 113.160.132.26:8080 | ✓ 1793ms | ✓ 1299ms | ✓ 1692ms | ✓ 1470ms | 否 | http |
| 194.67.99.223:1080 | ✓ 1227ms | ✓ 1921ms | ✓ 1691ms | 否 | 否 | http |
| 183.249.5.117:22222 | ✓ 724ms | ✓ 1157ms | ✓ 738ms | ✓ 996ms | ✓ 707ms | http |
| 167.103.31.122:8800 | ✓ 1370ms | 否 | ✓ 1407ms | 否 | ✓ 1544ms | http |
| 120.92.212.16:8890 | ✓ 1214ms | 否 | ✓ 1622ms | ✓ 1456ms | ✓ 980ms | http |
| 120.92.212.16:7890 | ✓ 1985ms | 否 | 否 | ✓ 1193ms | ✓ 942ms | http |
| 45.136.131.54:8448 | ✓ 633ms | ✓ 1060ms | ✓ 143ms | ✓ 684ms | ✓ 555ms | http |
| 38.34.179.57:8453 | ✓ 328ms | ✓ 1068ms | ✓ 452ms | ✓ 710ms | ✓ 789ms | http |
| 38.34.179.88:8446 | ✓ 547ms | ✓ 619ms | ✓ 674ms | ✓ 693ms | ✓ 680ms | http |
| 38.34.179.98:8453 | ✓ 614ms | ✓ 1280ms | ✓ 974ms | 否 | 否 | http |
| 38.34.179.6:8449 | ✓ 1212ms | ✓ 995ms | ✓ 755ms | ✓ 774ms | 否 | http |
| 38.34.183.164:8453 | ✓ 1256ms | ✓ 1402ms | ✓ 181ms | ✓ 707ms | ✓ 679ms | http |
| 38.34.183.211:8453 | ✓ 1313ms | ✓ 760ms | ✓ 765ms | ✓ 691ms | ✓ 703ms | http |
| 38.34.179.50:8452 | ✓ 361ms | ✓ 1075ms | ✓ 714ms | ✓ 1548ms | ✓ 563ms | http |
| 38.34.183.233:8448 | ✓ 1034ms | ✓ 1097ms | ✓ 708ms | ✓ 684ms | ✓ 755ms | http |
| 38.145.220.33:8448 | ✓ 293ms | ✓ 966ms | ✓ 1094ms | ✓ 1507ms | ✓ 498ms | http |
| 210.223.44.230:3128 | ✓ 1434ms | ✓ 1108ms | ✓ 861ms | ✓ 927ms | 否 | http |
| 38.34.179.86:8452 | ✓ 1378ms | ✓ 1311ms | ✓ 402ms | 否 | 否 | http |
| 38.34.179.39:8452 | ✓ 452ms | ✓ 934ms | ✓ 341ms | 否 | 否 | http |
| 45.136.130.171:8445 | ✓ 433ms | ✓ 933ms | ✓ 754ms | 否 | 否 | http |
| 38.34.179.97:8448 | ✓ 836ms | ✓ 1211ms | ✓ 90ms | ✓ 678ms | 否 | http |
| 38.145.220.37:8449 | ✓ 464ms | ✓ 594ms | ✓ 865ms | ✓ 709ms | ✓ 557ms | http |
| 103.84.95.54:7890 | ✓ 884ms | 否 | 否 | ✓ 803ms | ✓ 636ms | http |
| 38.34.179.16:8451 | ✓ 1547ms | ✓ 1479ms | ✓ 112ms | ✓ 831ms | ✓ 1053ms | http |
| 101.43.127.100:8877 | ✓ 879ms | ✓ 1110ms | ✓ 939ms | ✓ 1043ms | ✓ 843ms | http |
| 147.161.239.240:8800 | ✓ 1187ms | ✓ 1738ms | ✓ 1683ms | 否 | ✓ 1608ms | http |
| 137.220.150.104:6005 | ✓ 1950ms | ✓ 1960ms | ✓ 1334ms | ✓ 1185ms | ✓ 1642ms | http |
| 137.220.151.110:6005 | 否 | 否 | ✓ 1380ms | ✓ 1957ms | ✓ 1365ms | http |
| 38.34.183.222:8453 | ✓ 164ms | ✓ 614ms | ✓ 103ms | ✓ 741ms | ✓ 995ms | http |
| 183.249.5.105:22222 | ✓ 731ms | ✓ 827ms | ✓ 714ms | ✓ 885ms | ✓ 700ms | http |
| 137.220.150.22:6005 | ✓ 1515ms | 否 | ✓ 877ms | ✓ 1675ms | 否 | http |
| 45.136.131.57:8450 | ✓ 1243ms | ✓ 719ms | ✓ 111ms | ✓ 690ms | ✓ 530ms | http |
| 38.145.208.247:8446 | ✓ 751ms | ✓ 1072ms | ✓ 164ms | ✓ 845ms | ✓ 544ms | http |
| 38.34.179.97:8450 | ✓ 179ms | ✓ 644ms | ✓ 672ms | ✓ 847ms | ✓ 500ms | http |
| 91.238.105.64:2024 | ✓ 825ms | ✓ 1975ms | ✓ 1062ms | ✓ 1725ms | ✓ 1375ms | http |
| 8.219.97.248:80 | ✓ 1388ms | 否 | ✓ 1591ms | 否 | ✓ 1613ms | http |
| 150.107.140.238:3128 | ✓ 1715ms | 否 | ✓ 1979ms | 否 | ✓ 1934ms | http |
| 38.145.220.11:8445 | ✓ 663ms | ✓ 640ms | ✓ 138ms | ✓ 676ms | ✓ 689ms | http |
| 38.34.183.234:8450 | ✓ 684ms | ✓ 614ms | ✓ 257ms | ✓ 1031ms | ✓ 1519ms | http |
| 38.34.183.130:8452 | ✓ 683ms | ✓ 645ms | ✓ 252ms | ✓ 961ms | ✓ 1375ms | http |
| 183.249.5.110:22222 | ✓ 765ms | ✓ 873ms | ✓ 697ms | ✓ 950ms | ✓ 694ms | http |
| 183.249.5.111:22222 | ✓ 895ms | ✓ 823ms | ✓ 869ms | ✓ 1011ms | ✓ 678ms | http |
| 45.149.92.147:5001 | ✓ 630ms | 否 | ✓ 617ms | ✓ 791ms | ✓ 622ms | http |
| 38.34.179.25:8444 | ✓ 1133ms | ✓ 1672ms | ✓ 94ms | ✓ 685ms | ✓ 820ms | http |
| 38.34.179.23:8444 | ✓ 1157ms | 否 | ✓ 204ms | ✓ 820ms | ✓ 1039ms | http |
| 222.184.48.252:22222 | ✓ 861ms | ✓ 1222ms | ✓ 910ms | ✓ 1146ms | ✓ 857ms | http |
| 38.34.179.173:8452 | ✓ 1381ms | ✓ 1358ms | 否 | ✓ 1366ms | ✓ 937ms | http |
| 137.184.1.87:3128 | ✓ 283ms | ✓ 760ms | ✓ 993ms | 否 | 否 | http |
| 223.16.170.103:80 | ✓ 871ms | 否 | ✓ 824ms | 否 | ✓ 1073ms | http |
| 38.34.179.51:8449 | ✓ 660ms | ✓ 850ms | ✓ 275ms | ✓ 846ms | ✓ 560ms | http |
| 38.34.179.61:8445 | ✓ 1572ms | ✓ 692ms | ✓ 98ms | ✓ 686ms | ✓ 1148ms | http |
| 222.184.48.251:22222 | ✓ 969ms | ✓ 1444ms | ✓ 1029ms | 否 | 否 | http |
| 38.34.179.48:8449 | ✓ 655ms | ✓ 1224ms | ✓ 194ms | ✓ 874ms | ✓ 899ms | http |
| 137.220.150.170:6005 | ✓ 1967ms | 否 | ✓ 1827ms | ✓ 1812ms | 否 | http |
| 38.34.179.26:8450 | ✓ 543ms | ✓ 729ms | ✓ 669ms | ✓ 677ms | ✓ 529ms | http |
| 38.34.179.186:8444 | ✓ 671ms | ✓ 1056ms | ✓ 1308ms | ✓ 689ms | ✓ 541ms | http |
| 38.145.220.8:8452 | ✓ 559ms | ✓ 783ms | ✓ 857ms | 否 | ✓ 513ms | http |
| 45.136.130.178:8449 | ✓ 1490ms | 否 | ✓ 1008ms | ✓ 655ms | ✓ 494ms | http |
| 38.34.179.106:8446 | ✓ 1421ms | ✓ 1779ms | ✓ 111ms | ✓ 709ms | ✓ 746ms | http |
| 149.28.157.177:1080 | ✓ 1028ms | 否 | ✓ 795ms | ✓ 1207ms | ✓ 1306ms | http |
| 150.241.77.172:1080 | ✓ 780ms | 否 | ✓ 1750ms | 否 | ✓ 1505ms | http |
| 38.145.208.185:8449 | ✓ 1142ms | ✓ 817ms | ✓ 1278ms | 否 | 否 | http |
| 180.103.19.91:1080 | ✓ 1790ms | ✓ 1916ms | ✓ 1194ms | ✓ 1760ms | ✓ 1311ms | http |
| 59.46.216.131:30001 | ✓ 949ms | ✓ 1326ms | ✓ 1190ms | ✓ 1303ms | 否 | http |
| 38.34.179.57:8448 | ✓ 937ms | ✓ 894ms | ✓ 228ms | ✓ 706ms | ✓ 676ms | http |
| 45.136.131.35:8452 | ✓ 760ms | 否 | ✓ 521ms | ✓ 1196ms | ✓ 1675ms | http |
| 134.209.153.66:3128 | ✓ 1168ms | 否 | ✓ 1556ms | ✓ 1566ms | ✓ 1331ms | http |
| 45.136.130.186:8451 | ✓ 1357ms | ✓ 764ms | ✓ 445ms | ✓ 1408ms | ✓ 1880ms | http |
| 217.76.245.80:999 | ✓ 919ms | ✓ 1315ms | ✓ 1077ms | ✓ 1814ms | ✓ 1440ms | http |
| 38.34.179.96:8451 | ✓ 159ms | ✓ 617ms | ✓ 231ms | ✓ 777ms | ✓ 685ms | http |
| 38.34.179.165:8446 | ✓ 1594ms | ✓ 1579ms | 否 | ✓ 1216ms | ✓ 988ms | http |

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
