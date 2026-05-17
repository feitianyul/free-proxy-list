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

最后更新：2026-05-17 12:04:09 UTC（2026-05-17 20:04:09 UTC+8）

**代理总数：71**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 864ms | ✓ 984ms | ✓ 818ms | ✓ 1052ms | ✓ 822ms | http |
| 51.161.50.166:3128 | ✓ 438ms | ✓ 1710ms | ✓ 1249ms | 否 | ✓ 1287ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1431ms | ✓ 1307ms | 否 | ✓ 1111ms | http |
| 114.214.165.78:10810 | ✓ 1797ms | ✓ 1606ms | ✓ 1461ms | ✓ 1302ms | ✓ 1735ms | http |
| 185.40.77.94:1080 | ✓ 985ms | 否 | ✓ 1898ms | 否 | ✓ 1904ms | http |
| 185.200.188.234:10001 | ✓ 1436ms | 否 | ✓ 1883ms | 否 | ✓ 1660ms | http |
| 212.58.132.5:8888 | ✓ 1883ms | 否 | ✓ 1495ms | ✓ 1564ms | ✓ 1431ms | http |
| 115.231.181.40:8128 | ✓ 877ms | ✓ 1202ms | ✓ 910ms | ✓ 1218ms | ✓ 1478ms | http |
| 65.109.125.111:8443 | ✓ 842ms | 否 | ✓ 1996ms | 否 | ✓ 1688ms | http |
| 86.104.72.220:1081 | ✓ 621ms | ✓ 1160ms | ✓ 554ms | ✓ 1339ms | 否 | http |
| 91.242.229.129:8092 | ✓ 1445ms | 否 | ✓ 1029ms | ✓ 1814ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1876ms | ✓ 1387ms | 否 | 否 | ✓ 1016ms | http |
| 45.125.67.37:8443 | ✓ 877ms | 否 | ✓ 1652ms | ✓ 1140ms | ✓ 826ms | http |
| 128.199.116.219:9090 | ✓ 1480ms | 否 | ✓ 1249ms | ✓ 1028ms | ✓ 854ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 688ms | ✓ 1838ms | ✓ 1918ms | http |
| 148.230.4.241:999 | ✓ 642ms | 否 | ✓ 481ms | ✓ 1431ms | ✓ 1238ms | http |
| 170.106.136.181:31002 | ✓ 581ms | ✓ 744ms | ✓ 368ms | ✓ 605ms | ✓ 1058ms | http |
| 43.156.90.221:10808 | ✓ 692ms | 否 | ✓ 1948ms | ✓ 961ms | ✓ 1287ms | http |
| 128.199.114.189:9090 | ✓ 704ms | 否 | ✓ 1123ms | ✓ 1028ms | ✓ 784ms | http |
| 152.42.170.187:9090 | ✓ 685ms | 否 | ✓ 1202ms | ✓ 1000ms | ✓ 888ms | http |
| 146.190.80.158:9090 | ✓ 784ms | 否 | ✓ 1082ms | ✓ 1044ms | ✓ 881ms | http |
| 128.199.113.85:9090 | ✓ 944ms | 否 | ✓ 1261ms | ✓ 1055ms | ✓ 793ms | http |
| 121.230.9.96:1080 | ✓ 1339ms | 否 | ✓ 1350ms | ✓ 1485ms | ✓ 1093ms | http |
| 121.130.199.80:24003 | ✓ 1700ms | 否 | ✓ 1596ms | ✓ 1294ms | ✓ 901ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1421ms | ✓ 1408ms | ✓ 919ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 1381ms | ✓ 1029ms | ✓ 1786ms | http |
| 3.15.187.17:1080 | ✓ 1112ms | 否 | 否 | ✓ 1858ms | ✓ 1855ms | http |
| 190.12.150.244:999 | ✓ 1114ms | 否 | ✓ 1078ms | ✓ 1812ms | 否 | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1586ms | ✓ 1727ms | ✓ 1814ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1161ms | 否 | ✓ 1329ms | ✓ 950ms | http |
| 84.47.150.125:1080 | ✓ 1310ms | ✓ 1974ms | ✓ 1468ms | 否 | ✓ 1783ms | http |
| 116.171.106.78:3443 | 否 | 否 | ✓ 1845ms | ✓ 1684ms | ✓ 1734ms | http |
| 192.210.140.36:7913 | ✓ 1022ms | ✓ 1079ms | ✓ 563ms | ✓ 1131ms | ✓ 847ms | http |
| 150.107.140.238:3128 | ✓ 1435ms | 否 | ✓ 1913ms | ✓ 1084ms | ✓ 857ms | http |
| 129.80.238.83:444 | ✓ 424ms | ✓ 1151ms | ✓ 1351ms | 否 | 否 | http |
| 129.80.217.21:444 | ✓ 789ms | ✓ 1181ms | ✓ 393ms | ✓ 1193ms | ✓ 927ms | http |
| 114.214.170.41:27890 | ✓ 997ms | ✓ 1291ms | ✓ 1025ms | ✓ 1501ms | ✓ 1043ms | http |
| 166.88.55.83:7890 | ✓ 1554ms | ✓ 1162ms | ✓ 613ms | ✓ 769ms | ✓ 614ms | http |
| 133.18.123.225:26021 | ✓ 1476ms | 否 | ✓ 959ms | ✓ 1057ms | ✓ 1041ms | http |
| 103.21.220.141:3128 | ✓ 1595ms | 否 | ✓ 1702ms | ✓ 933ms | ✓ 691ms | http |
| 106.10.55.212:1121 | ✓ 1530ms | ✓ 1651ms | 否 | ✓ 1336ms | ✓ 886ms | http |
| 112.163.160.93:3128 | 否 | ✓ 1969ms | ✓ 797ms | ✓ 1848ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1948ms | ✓ 1106ms | ✓ 1920ms | 否 | 否 | http |
| 64.181.240.152:3128 | ✓ 1194ms | ✓ 628ms | ✓ 1235ms | 否 | 否 | http |
| 5.252.33.13:2025 | ✓ 1793ms | 否 | ✓ 1554ms | 否 | ✓ 1965ms | http |
| 128.199.254.13:9090 | ✓ 1384ms | 否 | ✓ 1185ms | ✓ 1021ms | 否 | http |
| 128.199.121.61:9090 | ✓ 1341ms | 否 | ✓ 814ms | ✓ 1024ms | ✓ 841ms | http |
| 104.248.151.93:9090 | ✓ 701ms | 否 | ✓ 703ms | 否 | ✓ 825ms | http |
| 159.223.41.216:9090 | ✓ 682ms | 否 | ✓ 696ms | ✓ 1028ms | ✓ 803ms | http |
| 121.130.177.28:8888 | ✓ 861ms | 否 | ✓ 1802ms | ✓ 1400ms | ✓ 1243ms | http |
| 57.129.144.178:40000 | ✓ 848ms | 否 | ✓ 734ms | 否 | ✓ 1511ms | http |
| 103.147.152.12:1095 | ✓ 1008ms | ✓ 1708ms | ✓ 1180ms | ✓ 1777ms | ✓ 1364ms | http |
| 8.154.21.175:3128 | ✓ 1239ms | ✓ 984ms | 否 | ✓ 1067ms | ✓ 1312ms | http |
| 3.101.133.120:80 | ✓ 433ms | 否 | ✓ 1583ms | ✓ 1128ms | ✓ 817ms | http |
| 45.174.77.1:999 | 否 | ✓ 1472ms | ✓ 1906ms | ✓ 1182ms | ✓ 865ms | http |
| 168.110.52.228:3128 | ✓ 1285ms | 否 | 否 | ✓ 1178ms | ✓ 603ms | http |
| 31.172.78.12:3128 | ✓ 841ms | ✓ 1884ms | ✓ 1084ms | 否 | ✓ 1790ms | http |
| 137.59.47.73:3128 | 否 | ✓ 1603ms | ✓ 1206ms | ✓ 1388ms | ✓ 1042ms | http |
| 103.147.152.12:1080 | 否 | 否 | ✓ 1194ms | ✓ 1748ms | ✓ 1377ms | http |
| 120.92.212.16:8890 | ✓ 883ms | ✓ 1228ms | ✓ 932ms | 否 | ✓ 922ms | http |
| 62.113.119.14:8080 | ✓ 1816ms | 否 | ✓ 1941ms | 否 | ✓ 1952ms | http |
| 138.197.68.35:4857 | ✓ 1127ms | 否 | ✓ 347ms | ✓ 1387ms | 否 | http |
| 185.71.196.92:1080 | ✓ 1065ms | 否 | ✓ 1247ms | 否 | ✓ 1913ms | http |
| 194.59.247.34:10808 | ✓ 1179ms | ✓ 1805ms | ✓ 1679ms | ✓ 1828ms | ✓ 1955ms | http |
| 45.174.175.26:999 | ✓ 831ms | ✓ 1686ms | ✓ 1265ms | 否 | 否 | http |
| 129.212.224.122:3128 | 否 | 否 | ✓ 1002ms | ✓ 1043ms | ✓ 803ms | http |
| 86.104.72.219:1081 | ✓ 324ms | 否 | ✓ 573ms | ✓ 1203ms | 否 | http |
| 86.104.72.219:1082 | ✓ 325ms | 否 | ✓ 569ms | ✓ 1224ms | 否 | http |
| 61.52.131.172:8443 | ✓ 900ms | ✓ 1114ms | ✓ 979ms | ✓ 1237ms | ✓ 928ms | http |
| 112.198.179.39:8000 | ✓ 1323ms | 否 | 否 | ✓ 1526ms | ✓ 1529ms | http |
| 121.130.199.80:24007 | 否 | ✓ 1197ms | ✓ 1776ms | ✓ 1277ms | ✓ 916ms | http |

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
