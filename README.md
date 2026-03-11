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

最后更新：2026-03-11 10:43:13 UTC（2026-03-11 18:43:13 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.63:8443 | ✓ 890ms | ✓ 832ms | ✓ 1147ms | ✓ 674ms | ✓ 505ms | http |
| 1.231.81.166:3128 | ✓ 709ms | ✓ 1426ms | ✓ 1811ms | ✓ 863ms | ✓ 772ms | http |
| 103.113.70.189:1081 | ✓ 380ms | ✓ 1179ms | 否 | 否 | ✓ 981ms | http |
| 103.84.95.54:7890 | ✓ 631ms | 否 | ✓ 1567ms | ✓ 1297ms | ✓ 614ms | http |
| 121.126.185.63:25152 | ✓ 1846ms | 否 | ✓ 1683ms | ✓ 1838ms | 否 | http |
| 14.143.222.113:10155 | ✓ 1680ms | 否 | ✓ 1082ms | ✓ 1350ms | 否 | http |
| 45.136.131.47:8443 | ✓ 379ms | ✓ 601ms | ✓ 84ms | ✓ 888ms | ✓ 683ms | http |
| 45.136.130.175:8443 | ✓ 380ms | ✓ 682ms | ✓ 83ms | ✓ 1666ms | ✓ 502ms | http |
| 45.136.130.188:8443 | ✓ 380ms | ✓ 678ms | ✓ 87ms | ✓ 1796ms | ✓ 1481ms | http |
| 45.136.130.191:8443 | ✓ 380ms | ✓ 1641ms | ✓ 73ms | 否 | ✓ 1096ms | http |
| 158.69.185.37:3129 | ✓ 491ms | ✓ 1656ms | ✓ 939ms | ✓ 1556ms | ✓ 1196ms | http |
| 160.238.65.5:3128 | ✓ 1172ms | ✓ 1699ms | ✓ 607ms | ✓ 1514ms | ✓ 1321ms | http |
| 160.238.65.3:3128 | ✓ 1173ms | 否 | ✓ 679ms | ✓ 1479ms | ✓ 1187ms | http |
| 160.238.65.7:3128 | ✓ 1165ms | ✓ 1756ms | ✓ 678ms | ✓ 1532ms | ✓ 1220ms | http |
| 160.238.65.2:3128 | ✓ 1166ms | 否 | ✓ 664ms | ✓ 1509ms | ✓ 1251ms | http |
| 160.238.65.9:3128 | ✓ 1439ms | ✓ 1700ms | ✓ 700ms | ✓ 1584ms | ✓ 1175ms | http |
| 160.238.65.4:3128 | ✓ 1173ms | 否 | ✓ 677ms | ✓ 1575ms | ✓ 1216ms | http |
| 160.238.65.6:3128 | ✓ 1172ms | ✓ 1698ms | ✓ 608ms | ✓ 1538ms | ✓ 1627ms | http |
| 160.238.65.8:3128 | ✓ 1439ms | 否 | ✓ 605ms | ✓ 1629ms | ✓ 1189ms | http |
| 115.231.181.40:8128 | ✓ 815ms | ✓ 1022ms | ✓ 1827ms | ✓ 1084ms | ✓ 654ms | http |
| 190.9.109.198:999 | ✓ 754ms | ✓ 1552ms | ✓ 1312ms | ✓ 1595ms | ✓ 1177ms | http |
| 47.77.193.180:1080 | ✓ 657ms | ✓ 1088ms | ✓ 623ms | ✓ 766ms | ✓ 497ms | http |
| 101.43.255.96:80 | ✓ 777ms | 否 | ✓ 818ms | ✓ 1057ms | ✓ 753ms | http |
| 39.104.201.40:7890 | ✓ 1831ms | ✓ 902ms | ✓ 1532ms | ✓ 1021ms | ✓ 1452ms | http |
| 152.70.98.46:8888 | ✓ 1417ms | 否 | 否 | ✓ 1279ms | ✓ 761ms | http |
| 81.70.169.194:80 | ✓ 1006ms | ✓ 1592ms | 否 | ✓ 1036ms | ✓ 847ms | http |
| 202.155.12.161:443 | ✓ 1298ms | 否 | 否 | ✓ 1670ms | ✓ 836ms | http |
| 120.92.212.16:7890 | ✓ 1306ms | 否 | ✓ 809ms | ✓ 1004ms | 否 | http |
| 62.113.119.14:8080 | ✓ 837ms | ✓ 1777ms | ✓ 1134ms | 否 | 否 | http |
| 43.162.113.116:3128 | ✓ 244ms | ✓ 650ms | ✓ 481ms | ✓ 599ms | ✓ 446ms | http |
| 194.213.18.200:443 | ✓ 1786ms | 否 | ✓ 1975ms | ✓ 1885ms | 否 | http |
| 14.225.212.37:7890 | ✓ 1704ms | 否 | ✓ 1605ms | ✓ 1426ms | ✓ 1004ms | http |
| 46.183.25.8:443 | ✓ 938ms | 否 | ✓ 216ms | ✓ 780ms | ✓ 633ms | http |
| 47.74.226.8:5001 | 否 | ✓ 1455ms | ✓ 1115ms | ✓ 1386ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1601ms | ✓ 1426ms | ✓ 639ms | ✓ 855ms | ✓ 699ms | http |
| 165.227.5.10:8888 | ✓ 1185ms | 否 | ✓ 1817ms | ✓ 701ms | 否 | http |
| 24.199.124.152:3128 | ✓ 523ms | ✓ 1099ms | ✓ 1098ms | ✓ 676ms | ✓ 473ms | http |
| 27.254.99.183:8118 | ✓ 1254ms | 否 | ✓ 1118ms | ✓ 1249ms | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1298ms | ✓ 327ms | 否 | ✓ 924ms | http |
| 91.107.141.42:8081 | ✓ 719ms | 否 | ✓ 1946ms | ✓ 1893ms | 否 | http |
| 200.174.198.32:8888 | ✓ 1035ms | 否 | ✓ 1855ms | 否 | ✓ 1884ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1132ms | ✓ 790ms | ✓ 980ms | 否 | http |
| 107.172.125.217:3128 | ✓ 588ms | 否 | ✓ 564ms | ✓ 662ms | ✓ 738ms | http |
| 205.209.118.30:3138 | ✓ 592ms | 否 | ✓ 967ms | ✓ 1411ms | ✓ 1074ms | http |
| 42.96.16.158:1311 | ✓ 1348ms | 否 | 否 | ✓ 1176ms | ✓ 954ms | http |
| 116.80.47.62:3172 | ✓ 1744ms | 否 | ✓ 1517ms | 否 | ✓ 1799ms | http |
| 116.80.49.159:3172 | ✓ 1744ms | 否 | ✓ 1520ms | ✓ 1790ms | 否 | http |
| 45.140.147.82:1081 | ✓ 1159ms | ✓ 1755ms | ✓ 1448ms | ✓ 1790ms | ✓ 1247ms | http |
| 190.212.131.238:3128 | ✓ 771ms | ✓ 1553ms | ✓ 737ms | 否 | 否 | http |
| 162.248.165.72:1080 | ✓ 1492ms | 否 | 否 | ✓ 1786ms | ✓ 1917ms | http |
| 221.202.27.194:10810 | 否 | ✓ 1759ms | 否 | ✓ 1428ms | ✓ 1836ms | http |
| 116.80.49.163:3172 | ✓ 1661ms | 否 | ✓ 1547ms | 否 | ✓ 1626ms | http |
| 113.177.131.2:3128 | ✓ 1946ms | ✓ 1926ms | 否 | 否 | ✓ 1004ms | http |
| 103.183.10.169:3125 | ✓ 1880ms | 否 | 否 | ✓ 1440ms | ✓ 1365ms | http |
| 34.101.184.164:3128 | ✓ 1581ms | 否 | ✓ 1381ms | ✓ 1496ms | 否 | http |
| 116.80.49.170:3172 | ✓ 1454ms | 否 | ✓ 1531ms | 否 | ✓ 1599ms | http |
| 116.80.49.172:3172 | ✓ 1454ms | 否 | ✓ 1528ms | ✓ 1778ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1633ms | ✓ 1818ms | ✓ 1725ms | 否 | ✓ 1553ms | http |
| 103.183.10.203:3125 | ✓ 1909ms | 否 | ✓ 1750ms | ✓ 1850ms | 否 | http |
| 66.228.47.125:110 | ✓ 906ms | 否 | 否 | ✓ 1592ms | ✓ 1367ms | http |
| 178.236.245.59:3128 | ✓ 1060ms | 否 | ✓ 1785ms | 否 | ✓ 1934ms | http |
| 45.136.130.223:8443 | ✓ 210ms | ✓ 1917ms | ✓ 1624ms | ✓ 678ms | ✓ 517ms | http |
| 178.236.245.17:3128 | ✓ 1184ms | ✓ 1972ms | ✓ 877ms | ✓ 1908ms | ✓ 1709ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1176ms | ✓ 1266ms | ✓ 1515ms | http |
| 116.80.49.167:3172 | ✓ 1557ms | 否 | 否 | ✓ 1805ms | ✓ 1645ms | http |
| 116.80.49.166:3172 | 否 | 否 | ✓ 1949ms | ✓ 1802ms | ✓ 1612ms | http |
| 168.235.110.63:3128 | ✓ 620ms | 否 | ✓ 1006ms | ✓ 1288ms | ✓ 990ms | http |
| 103.39.51.190:8080 | ✓ 1740ms | 否 | 否 | ✓ 1714ms | ✓ 1532ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1583ms | ✓ 1669ms | ✓ 1150ms | 否 | http |
| 106.14.203.63:3333 | ✓ 602ms | ✓ 1873ms | ✓ 769ms | ✓ 831ms | ✓ 718ms | http |
| 61.52.131.172:8443 | ✓ 725ms | ✓ 861ms | ✓ 736ms | ✓ 926ms | ✓ 679ms | http |
| 123.57.0.163:8888 | 否 | 否 | ✓ 1543ms | ✓ 1262ms | ✓ 1162ms | http |
| 211.171.114.154:3128 | ✓ 1467ms | 否 | 否 | ✓ 1750ms | ✓ 1092ms | http |
| 121.138.61.193:8118 | 否 | 否 | ✓ 1378ms | ✓ 1009ms | ✓ 814ms | http |
| 95.3.9.78:3128 | ✓ 1563ms | 否 | 否 | ✓ 1853ms | ✓ 1438ms | http |
| 14.225.222.213:7890 | ✓ 1525ms | ✓ 1987ms | ✓ 1030ms | 否 | 否 | http |
| 8.219.97.248:80 | ✓ 1747ms | 否 | ✓ 1327ms | 否 | ✓ 1622ms | http |
| 152.42.213.210:8080 | ✓ 1199ms | 否 | ✓ 1223ms | ✓ 1152ms | ✓ 843ms | http |
| 61.155.242.150:5566 | ✓ 1410ms | ✓ 1704ms | 否 | ✓ 1906ms | ✓ 920ms | http |
| 95.3.9.78:8080 | ✓ 862ms | 否 | 否 | ✓ 1816ms | ✓ 1429ms | http |

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
