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

最后更新：2026-03-23 03:33:16 UTC（2026-03-23 11:33:16 UTC+8）

**代理总数：296**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 295 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 296 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 803ms | 否 | ✓ 771ms | ✓ 980ms | ✓ 786ms | http |
| 147.161.210.140:8800 | 否 | 否 | ✓ 848ms | ✓ 1111ms | ✓ 1787ms | http |
| 101.47.73.135:3128 | ✓ 1114ms | 否 | ✓ 1633ms | ✓ 1744ms | ✓ 1403ms | http |
| 35.225.22.61:80 | ✓ 746ms | ✓ 1383ms | 否 | ✓ 1187ms | ✓ 1004ms | http |
| 38.34.179.162:8451 | ✓ 789ms | ✓ 1076ms | ✓ 595ms | ✓ 880ms | ✓ 670ms | http |
| 45.136.130.172:8443 | ✓ 757ms | ✓ 1030ms | ✓ 775ms | ✓ 993ms | ✓ 681ms | http |
| 45.136.130.175:8443 | ✓ 725ms | ✓ 1028ms | ✓ 781ms | ✓ 1000ms | ✓ 680ms | http |
| 45.136.130.174:8447 | ✓ 717ms | ✓ 1706ms | ✓ 225ms | ✓ 957ms | ✓ 667ms | http |
| 38.34.179.150:8449 | ✓ 739ms | ✓ 1229ms | ✓ 614ms | ✓ 860ms | ✓ 656ms | http |
| 45.136.130.173:8447 | ✓ 726ms | 否 | ✓ 244ms | ✓ 957ms | ✓ 670ms | http |
| 38.34.179.40:8446 | ✓ 770ms | ✓ 932ms | ✓ 844ms | ✓ 1044ms | ✓ 868ms | http |
| 20.78.118.91:8561 | ✓ 589ms | ✓ 1252ms | ✓ 776ms | ✓ 1010ms | ✓ 743ms | http |
| 20.27.11.248:8561 | ✓ 629ms | ✓ 1184ms | ✓ 811ms | ✓ 967ms | ✓ 770ms | http |
| 20.27.14.220:8561 | ✓ 643ms | ✓ 1106ms | ✓ 955ms | ✓ 974ms | ✓ 731ms | http |
| 20.27.13.35:8561 | ✓ 634ms | ✓ 970ms | ✓ 1019ms | ✓ 967ms | ✓ 775ms | http |
| 20.27.15.111:8561 | ✓ 621ms | ✓ 1189ms | ✓ 798ms | ✓ 963ms | ✓ 776ms | http |
| 20.210.39.153:8561 | ✓ 594ms | 否 | ✓ 591ms | ✓ 961ms | ✓ 734ms | http |
| 20.78.26.206:8561 | ✓ 597ms | 否 | ✓ 573ms | ✓ 978ms | ✓ 736ms | http |
| 38.34.179.83:8448 | ✓ 1522ms | ✓ 1967ms | ✓ 351ms | ✓ 1030ms | ✓ 1705ms | http |
| 38.34.179.76:8452 | ✓ 788ms | 否 | ✓ 878ms | ✓ 876ms | ✓ 900ms | http |
| 38.34.179.80:8452 | ✓ 1017ms | 否 | ✓ 643ms | ✓ 870ms | ✓ 912ms | http |
| 219.117.204.211:7799 | ✓ 970ms | ✓ 1577ms | ✓ 1948ms | ✓ 1282ms | ✓ 801ms | http |
| 38.34.183.8:8450 | ✓ 1230ms | ✓ 1959ms | ✓ 616ms | 否 | ✓ 1429ms | http |
| 217.174.244.117:3129 | ✓ 1768ms | ✓ 1535ms | ✓ 1359ms | ✓ 1974ms | ✓ 1554ms | http |
| 222.184.48.251:22222 | ✓ 1076ms | 否 | ✓ 1059ms | ✓ 1728ms | ✓ 1047ms | http |
| 38.34.179.26:8450 | ✓ 1821ms | ✓ 1115ms | 否 | ✓ 1447ms | ✓ 755ms | http |
| 137.220.151.110:6005 | 否 | 否 | ✓ 1063ms | ✓ 1401ms | ✓ 1395ms | http |
| 167.103.34.108:8800 | 否 | 否 | ✓ 1243ms | ✓ 1482ms | ✓ 1371ms | http |
| 137.220.150.22:6005 | 否 | 否 | ✓ 939ms | ✓ 1317ms | ✓ 1289ms | http |
| 113.160.132.26:8080 | ✓ 1999ms | 否 | ✓ 1068ms | ✓ 1435ms | 否 | http |
| 38.34.179.60:8450 | 否 | ✓ 1459ms | ✓ 366ms | ✓ 1998ms | ✓ 1495ms | http |
| 137.220.150.152:6005 | 否 | 否 | ✓ 1882ms | ✓ 1938ms | ✓ 1822ms | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1687ms | ✓ 1429ms | ✓ 1066ms | http |
| 45.136.130.186:8451 | ✓ 1247ms | ✓ 1471ms | ✓ 1461ms | ✓ 1262ms | ✓ 1262ms | http |
| 45.136.130.178:8449 | ✓ 1425ms | 否 | ✓ 1141ms | ✓ 1078ms | ✓ 1819ms | http |
| 147.161.239.240:8800 | ✓ 545ms | 否 | ✓ 1024ms | ✓ 1657ms | ✓ 1530ms | http |
| 167.103.31.122:8800 | ✓ 1353ms | 否 | ✓ 1313ms | ✓ 1631ms | 否 | http |
| 38.34.179.59:8444 | ✓ 1661ms | ✓ 892ms | ✓ 1135ms | 否 | ✓ 661ms | http |
| 115.231.181.40:8128 | ✓ 1979ms | ✓ 1952ms | 否 | 否 | ✓ 1078ms | http |
| 194.67.99.223:1080 | ✓ 558ms | 否 | 否 | ✓ 1847ms | ✓ 1114ms | http |
| 183.249.5.117:22222 | ✓ 997ms | ✓ 1297ms | ✓ 1012ms | ✓ 1309ms | ✓ 970ms | http |
| 103.84.95.54:7890 | ✓ 739ms | 否 | ✓ 738ms | 否 | ✓ 930ms | http |
| 38.34.179.173:8452 | ✓ 854ms | ✓ 882ms | ✓ 399ms | ✓ 881ms | ✓ 751ms | http |
| 38.34.179.178:8445 | ✓ 831ms | ✓ 852ms | ✓ 429ms | ✓ 887ms | ✓ 749ms | http |
| 38.34.179.172:8451 | ✓ 853ms | ✓ 968ms | ✓ 309ms | ✓ 899ms | ✓ 736ms | http |
| 38.34.179.174:8453 | ✓ 823ms | ✓ 926ms | ✓ 368ms | ✓ 900ms | ✓ 738ms | http |
| 38.34.179.54:8446 | ✓ 852ms | ✓ 906ms | ✓ 384ms | ✓ 1052ms | ✓ 670ms | http |
| 38.34.179.48:8449 | ✓ 831ms | ✓ 868ms | ✓ 437ms | ✓ 1047ms | ✓ 657ms | http |
| 45.136.131.43:8448 | ✓ 841ms | ✓ 1265ms | ✓ 339ms | ✓ 917ms | ✓ 716ms | http |
| 45.136.131.59:8450 | ✓ 824ms | ✓ 1497ms | ✓ 248ms | ✓ 868ms | ✓ 682ms | http |
| 45.136.130.168:8452 | ✓ 830ms | ✓ 1483ms | ✓ 285ms | ✓ 887ms | ✓ 706ms | http |
| 45.136.131.38:8448 | ✓ 822ms | ✓ 1337ms | ✓ 294ms | ✓ 919ms | ✓ 695ms | http |
| 38.34.179.61:8445 | ✓ 819ms | ✓ 1646ms | ✓ 261ms | ✓ 928ms | ✓ 672ms | http |
| 38.34.178.155:8448 | ✓ 822ms | ✓ 1853ms | ✓ 1232ms | 否 | 否 | http |
| 38.34.179.81:8446 | ✓ 816ms | ✓ 1041ms | ✓ 444ms | ✓ 901ms | ✓ 695ms | http |
| 38.34.183.11:8446 | ✓ 834ms | ✓ 894ms | ✓ 389ms | ✓ 892ms | ✓ 968ms | http |
| 38.34.179.50:8452 | ✓ 1792ms | ✓ 853ms | ✓ 394ms | ✓ 1044ms | 否 | http |
| 38.34.179.228:8453 | ✓ 1064ms | ✓ 870ms | ✓ 520ms | ✓ 1570ms | 否 | http |
| 38.34.183.219:8446 | ✓ 839ms | ✓ 1942ms | ✓ 546ms | ✓ 890ms | 否 | http |
| 45.136.131.37:8448 | ✓ 809ms | ✓ 857ms | ✓ 764ms | ✓ 931ms | ✓ 702ms | http |
| 38.34.179.51:8449 | ✓ 810ms | ✓ 877ms | ✓ 430ms | ✓ 1065ms | ✓ 695ms | http |
| 45.136.130.169:8444 | ✓ 805ms | ✓ 1859ms | ✓ 286ms | ✓ 901ms | ✓ 663ms | http |
| 45.136.131.41:8448 | ✓ 818ms | ✓ 837ms | ✓ 756ms | ✓ 943ms | ✓ 732ms | http |
| 45.136.131.46:8448 | ✓ 825ms | ✓ 905ms | ✓ 702ms | ✓ 933ms | ✓ 714ms | http |
| 45.136.131.40:8448 | ✓ 803ms | ✓ 851ms | ✓ 767ms | ✓ 949ms | ✓ 728ms | http |
| 38.34.183.16:8452 | ✓ 812ms | ✓ 832ms | ✓ 458ms | ✓ 1637ms | ✓ 659ms | http |
| 45.136.130.195:8444 | ✓ 826ms | ✓ 845ms | ✓ 451ms | ✓ 1387ms | 否 | http |
| 45.136.130.171:8445 | ✓ 819ms | ✓ 885ms | ✓ 811ms | ✓ 908ms | ✓ 714ms | http |
| 45.136.130.167:8444 | ✓ 833ms | ✓ 1958ms | ✓ 264ms | ✓ 888ms | ✓ 1104ms | http |
| 38.34.179.53:8452 | ✓ 1808ms | ✓ 926ms | ✓ 429ms | ✓ 1105ms | 否 | http |
| 45.136.130.193:8444 | ✓ 842ms | ✓ 901ms | ✓ 725ms | ✓ 1830ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1457ms | ✓ 1594ms | ✓ 1270ms | 否 | 否 | http |
| 38.34.179.73:8446 | ✓ 803ms | ✓ 1034ms | ✓ 446ms | ✓ 863ms | ✓ 694ms | http |
| 38.34.179.202:8445 | ✓ 806ms | ✓ 1041ms | ✓ 637ms | ✓ 1829ms | 否 | http |
| 38.34.179.56:8452 | ✓ 1790ms | ✓ 821ms | ✓ 305ms | ✓ 1069ms | 否 | http |
| 45.136.130.197:8444 | ✓ 801ms | ✓ 1861ms | ✓ 1502ms | 否 | 否 | http |
| 38.34.183.47:8452 | ✓ 818ms | ✓ 879ms | ✓ 418ms | ✓ 879ms | ✓ 975ms | http |
| 38.34.179.11:8449 | ✓ 827ms | ✓ 1672ms | ✓ 264ms | ✓ 894ms | ✓ 667ms | http |
| 38.34.179.49:8450 | ✓ 828ms | ✓ 887ms | ✓ 420ms | ✓ 1037ms | ✓ 671ms | http |
| 45.136.131.36:8450 | ✓ 821ms | ✓ 892ms | ✓ 713ms | ✓ 931ms | ✓ 1134ms | http |
| 45.136.131.39:8443 | ✓ 805ms | ✓ 1852ms | ✓ 303ms | ✓ 1002ms | ✓ 688ms | http |
| 38.34.179.23:8444 | ✓ 302ms | ✓ 841ms | ✓ 264ms | 否 | 否 | http |
| 38.34.178.141:8453 | ✓ 318ms | ✓ 885ms | ✓ 441ms | 否 | 否 | http |
| 157.230.220.25:4857 | 否 | ✓ 1527ms | 否 | ✓ 1251ms | ✓ 767ms | http |
| 74.103.66.15:80 | ✓ 738ms | ✓ 1576ms | ✓ 1284ms | ✓ 1057ms | ✓ 794ms | http |
| 172.212.68.37:3128 | ✓ 706ms | 否 | ✓ 817ms | ✓ 1488ms | ✓ 1108ms | http |
| 34.150.20.6:8888 | ✓ 751ms | ✓ 1567ms | ✓ 1074ms | ✓ 937ms | ✓ 749ms | http |
| 38.34.179.16:8451 | ✓ 731ms | ✓ 1874ms | ✓ 981ms | ✓ 897ms | ✓ 1891ms | http |
| 106.14.203.63:3333 | ✓ 929ms | ✓ 1185ms | ✓ 1269ms | ✓ 1319ms | ✓ 960ms | http |
| 1.231.81.166:3128 | ✓ 1562ms | 否 | ✓ 1648ms | ✓ 1407ms | ✓ 1251ms | http |
| 120.92.212.16:7890 | ✓ 1148ms | ✓ 1361ms | ✓ 1110ms | 否 | ✓ 1062ms | http |
| 101.43.127.100:8877 | ✓ 1145ms | 否 | ✓ 1037ms | 否 | ✓ 1042ms | http |
| 64.227.76.27:1080 | ✓ 1736ms | 否 | ✓ 1696ms | 否 | ✓ 1958ms | http |
| 181.78.44.63:999 | ✓ 1029ms | 否 | ✓ 1515ms | 否 | ✓ 1679ms | http |
| 218.89.134.230:3333 | ✓ 1577ms | ✓ 1723ms | 否 | ✓ 1746ms | ✓ 1369ms | http |
| 62.171.161.88:2018 | ✓ 449ms | ✓ 1572ms | ✓ 980ms | ✓ 1801ms | ✓ 1182ms | http |
| 104.168.158.236:10808 | ✓ 427ms | 否 | ✓ 1897ms | ✓ 1252ms | ✓ 910ms | http |
| 38.145.208.244:8448 | ✓ 1120ms | ✓ 1048ms | ✓ 603ms | ✓ 1299ms | ✓ 1091ms | http |
| 38.34.179.25:8444 | ✓ 743ms | 否 | ✓ 789ms | ✓ 893ms | ✓ 868ms | http |
| 183.249.5.111:22222 | ✓ 981ms | ✓ 1619ms | ✓ 849ms | ✓ 1165ms | ✓ 1002ms | http |

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
