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

最后更新：2026-03-11 21:35:31 UTC（2026-03-12 05:35:31 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.63:8443 | ✓ 306ms | ✓ 938ms | ✓ 299ms | ✓ 969ms | ✓ 730ms | http |
| 194.213.18.200:443 | ✓ 1941ms | ✓ 911ms | ✓ 130ms | ✓ 901ms | 否 | http |
| 1.231.81.166:3128 | ✓ 818ms | ✓ 1062ms | ✓ 775ms | ✓ 1156ms | ✓ 882ms | http |
| 202.155.12.161:443 | ✓ 1506ms | ✓ 1820ms | ✓ 1372ms | 否 | 否 | http |
| 171.251.172.78:5109 | ✓ 1483ms | 否 | ✓ 1482ms | ✓ 1775ms | ✓ 1576ms | http |
| 205.209.118.30:3138 | ✓ 1740ms | 否 | ✓ 1535ms | 否 | ✓ 1935ms | http |
| 45.136.130.175:8443 | ✓ 747ms | ✓ 913ms | ✓ 302ms | ✓ 934ms | ✓ 742ms | http |
| 45.136.130.223:8443 | ✓ 741ms | ✓ 885ms | ✓ 364ms | ✓ 932ms | ✓ 739ms | http |
| 45.136.131.47:8443 | ✓ 742ms | ✓ 881ms | ✓ 304ms | ✓ 947ms | ✓ 786ms | http |
| 45.136.130.188:8443 | ✓ 1547ms | ✓ 1029ms | ✓ 523ms | ✓ 962ms | ✓ 729ms | http |
| 152.42.213.210:8080 | ✓ 961ms | 否 | ✓ 1634ms | ✓ 1310ms | ✓ 1313ms | http |
| 165.227.5.10:8888 | ✓ 1601ms | ✓ 1833ms | 否 | ✓ 1602ms | 否 | http |
| 103.84.95.54:7890 | ✓ 902ms | 否 | ✓ 1429ms | ✓ 1161ms | ✓ 973ms | http |
| 111.48.191.1:7890 | ✓ 886ms | ✓ 1159ms | ✓ 915ms | 否 | 否 | http |
| 137.184.1.87:3128 | ✓ 489ms | ✓ 939ms | ✓ 441ms | ✓ 996ms | ✓ 772ms | http |
| 103.35.188.243:3128 | ✓ 61ms | ✓ 1696ms | 否 | ✓ 1044ms | ✓ 808ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1117ms | 否 | ✓ 950ms | ✓ 1203ms | http |
| 81.70.169.194:80 | ✓ 1116ms | ✓ 1534ms | ✓ 1166ms | ✓ 1569ms | ✓ 1223ms | http |
| 101.43.255.96:80 | ✓ 1210ms | ✓ 1408ms | ✓ 1210ms | ✓ 1579ms | ✓ 1228ms | http |
| 115.231.181.40:8128 | ✓ 1051ms | 否 | ✓ 1129ms | 否 | ✓ 1110ms | http |
| 138.124.53.25:7443 | ✓ 381ms | 否 | ✓ 1772ms | ✓ 1598ms | ✓ 1623ms | http |
| 35.225.22.61:80 | 否 | ✓ 1191ms | ✓ 349ms | ✓ 1051ms | ✓ 883ms | http |
| 107.173.52.58:7890 | ✓ 666ms | ✓ 1317ms | 否 | ✓ 1224ms | ✓ 1182ms | http |
| 113.160.132.26:8080 | ✓ 1099ms | ✓ 1553ms | ✓ 1904ms | ✓ 1459ms | ✓ 1179ms | http |
| 45.186.6.104:3128 | ✓ 1176ms | ✓ 1821ms | ✓ 1528ms | 否 | 否 | http |
| 46.183.25.8:443 | ✓ 1666ms | 否 | ✓ 1082ms | ✓ 1259ms | ✓ 1558ms | http |
| 39.104.201.40:7890 | ✓ 1393ms | ✓ 1406ms | 否 | ✓ 1443ms | ✓ 1112ms | http |
| 168.235.110.63:3128 | ✓ 157ms | 否 | ✓ 1006ms | ✓ 892ms | ✓ 728ms | http |
| 83.219.250.8:62920 | ✓ 1596ms | 否 | ✓ 1600ms | 否 | ✓ 1381ms | http |
| 101.47.73.135:3128 | ✓ 1493ms | 否 | 否 | ✓ 1495ms | ✓ 1953ms | http |
| 190.212.131.238:3128 | ✓ 1645ms | 否 | ✓ 1814ms | 否 | ✓ 1755ms | http |
| 45.136.130.191:8443 | ✓ 685ms | ✓ 1033ms | ✓ 911ms | ✓ 1091ms | ✓ 718ms | http |
| 43.167.227.161:1080 | 否 | ✓ 1286ms | ✓ 1051ms | 否 | ✓ 1239ms | http |
| 45.168.238.193:8443 | ✓ 599ms | ✓ 1317ms | ✓ 698ms | ✓ 1230ms | ✓ 991ms | http |
| 103.82.23.118:5221 | 否 | 否 | ✓ 1880ms | ✓ 1979ms | ✓ 1865ms | http |
| 47.77.193.180:1080 | ✓ 869ms | ✓ 978ms | ✓ 398ms | ✓ 878ms | ✓ 686ms | http |
| 18.192.100.176:3282 | ✓ 1047ms | 否 | ✓ 1751ms | 否 | ✓ 1948ms | http |
| 121.230.9.241:1080 | ✓ 1245ms | ✓ 1500ms | ✓ 1184ms | 否 | 否 | http |
| 116.236.189.93:29999 | ✓ 965ms | 否 | ✓ 1299ms | 否 | ✓ 1777ms | http |
| 152.42.213.210:443 | 否 | 否 | ✓ 1595ms | ✓ 1354ms | ✓ 1225ms | http |
| 88.80.150.82:8080 | ✓ 1192ms | 否 | 否 | ✓ 1966ms | ✓ 1679ms | https |
| 91.107.141.42:8081 | 否 | 否 | ✓ 1877ms | ✓ 1557ms | ✓ 1417ms | http |
| 34.101.184.164:3128 | ✓ 1689ms | 否 | ✓ 1843ms | ✓ 1459ms | ✓ 1542ms | http |
| 86.53.183.16:1080 | ✓ 1338ms | 否 | ✓ 1705ms | ✓ 1822ms | ✓ 1521ms | http |
| 8.219.97.248:80 | ✓ 1776ms | 否 | 否 | ✓ 1750ms | ✓ 1759ms | http |
| 223.16.170.103:80 | ✓ 1107ms | ✓ 1928ms | ✓ 1906ms | ✓ 1313ms | ✓ 1410ms | http |
| 1.234.153.14:80 | ✓ 1386ms | ✓ 1329ms | ✓ 913ms | ✓ 1091ms | ✓ 876ms | http |
| 59.46.216.131:30001 | ✓ 1290ms | ✓ 1552ms | ✓ 1336ms | ✓ 1674ms | ✓ 1191ms | http |
| 45.136.198.40:3128 | ✓ 1115ms | ✓ 1834ms | ✓ 1596ms | 否 | ✓ 1850ms | http |
| 222.109.119.178:3128 | ✓ 986ms | ✓ 1281ms | ✓ 992ms | 否 | 否 | http |
| 121.138.61.193:8803 | ✓ 1339ms | 否 | ✓ 1222ms | ✓ 1283ms | ✓ 1103ms | http |
| 180.127.149.252:1080 | ✓ 1117ms | ✓ 1374ms | ✓ 1134ms | ✓ 1350ms | ✓ 1149ms | http |
| 220.197.44.36:3128 | ✓ 1345ms | ✓ 1674ms | ✓ 1372ms | 否 | ✓ 1521ms | http |
| 162.240.154.26:3128 | ✓ 1095ms | 否 | ✓ 1701ms | ✓ 1275ms | 否 | http |
| 103.113.70.189:1081 | ✓ 293ms | ✓ 949ms | 否 | ✓ 1460ms | ✓ 1076ms | http |
| 209.126.10.139:3128 | ✓ 596ms | ✓ 1066ms | ✓ 1142ms | ✓ 1153ms | ✓ 948ms | http |
| 201.144.20.238:3128 | ✓ 775ms | 否 | ✓ 1038ms | ✓ 1392ms | ✓ 953ms | http |
| 109.234.38.35:3128 | ✓ 1085ms | 否 | ✓ 1233ms | ✓ 1238ms | ✓ 1079ms | http |
| 223.16.170.103:3128 | ✓ 1763ms | ✓ 1879ms | ✓ 1689ms | ✓ 1631ms | ✓ 1334ms | http |
| 103.39.51.190:8080 | ✓ 1872ms | 否 | 否 | ✓ 1620ms | ✓ 1668ms | http |
| 211.171.114.154:3128 | ✓ 1721ms | 否 | ✓ 1884ms | 否 | ✓ 1811ms | http |
| 139.159.99.242:8080 | ✓ 1036ms | ✓ 1232ms | ✓ 993ms | ✓ 1265ms | ✓ 1542ms | http |
| 61.52.131.172:8443 | ✓ 1034ms | ✓ 1332ms | ✓ 1053ms | ✓ 1370ms | ✓ 1074ms | http |
| 164.92.148.68:3128 | ✓ 1535ms | ✓ 1805ms | ✓ 1881ms | ✓ 1616ms | ✓ 1564ms | http |
| 152.70.98.46:8888 | ✓ 1677ms | ✓ 1604ms | ✓ 1634ms | ✓ 1030ms | ✓ 904ms | http |

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
