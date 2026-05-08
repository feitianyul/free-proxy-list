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

最后更新：2026-05-08 12:12:31 UTC（2026-05-08 20:12:31 UTC+8）

**代理总数：61**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | 否 | ✓ 1801ms | ✓ 1540ms | ✓ 1406ms | ✓ 926ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1536ms | ✓ 1366ms | ✓ 1289ms | ✓ 1172ms | http |
| 115.231.181.40:8128 | ✓ 1957ms | 否 | ✓ 1037ms | ✓ 1614ms | ✓ 1695ms | http |
| 103.147.152.12:1080 | ✓ 640ms | 否 | ✓ 606ms | ✓ 1775ms | ✓ 1310ms | http |
| 8.219.97.248:80 | ✓ 1737ms | 否 | ✓ 1839ms | 否 | ✓ 1991ms | http |
| 45.125.67.37:8443 | ✓ 1008ms | 否 | ✓ 743ms | ✓ 1088ms | ✓ 1602ms | http |
| 212.224.88.212:443 | ✓ 581ms | 否 | ✓ 1197ms | 否 | ✓ 1705ms | http |
| 185.221.237.57:443 | ✓ 1609ms | 否 | ✓ 706ms | 否 | ✓ 1606ms | http |
| 185.221.237.57:8443 | ✓ 674ms | ✓ 1962ms | ✓ 1466ms | 否 | ✓ 1864ms | http |
| 65.109.125.111:8443 | ✓ 765ms | 否 | ✓ 1728ms | ✓ 1823ms | 否 | http |
| 84.47.150.125:1080 | ✓ 1086ms | 否 | ✓ 1601ms | 否 | ✓ 1909ms | http |
| 152.32.132.190:7890 | ✓ 1237ms | 否 | 否 | ✓ 1117ms | ✓ 1530ms | http |
| 62.133.60.126:24558 | ✓ 762ms | 否 | ✓ 1285ms | 否 | ✓ 1841ms | http |
| 91.242.229.129:8092 | ✓ 1682ms | ✓ 1632ms | ✓ 1697ms | 否 | ✓ 1719ms | http |
| 147.45.178.211:14658 | ✓ 1131ms | 否 | ✓ 1121ms | 否 | ✓ 1507ms | http |
| 45.153.231.229:8080 | ✓ 809ms | 否 | ✓ 724ms | 否 | ✓ 1871ms | http |
| 154.12.231.32:80 | ✓ 810ms | ✓ 1480ms | 否 | 否 | ✓ 1672ms | http |
| 79.137.205.44:40000 | ✓ 1100ms | 否 | ✓ 1341ms | ✓ 1990ms | 否 | http |
| 212.58.132.5:8888 | ✓ 995ms | 否 | ✓ 997ms | ✓ 1439ms | ✓ 1111ms | http |
| 59.46.216.131:30001 | ✓ 1198ms | 否 | ✓ 1123ms | ✓ 1804ms | ✓ 1206ms | http |
| 43.156.132.113:3128 | ✓ 1482ms | ✓ 1574ms | ✓ 767ms | ✓ 1042ms | ✓ 871ms | http |
| 148.230.4.241:999 | ✓ 751ms | ✓ 1870ms | ✓ 1665ms | ✓ 1450ms | ✓ 1134ms | http |
| 206.206.126.177:2412 | ✓ 1640ms | 否 | ✓ 1938ms | ✓ 1868ms | 否 | http |
| 137.59.47.73:3128 | ✓ 1122ms | ✓ 1693ms | ✓ 989ms | ✓ 1259ms | ✓ 1028ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 854ms | ✓ 1138ms | ✓ 879ms | http |
| 217.76.245.80:999 | ✓ 1101ms | ✓ 1584ms | ✓ 1317ms | ✓ 1394ms | ✓ 1183ms | http |
| 86.104.72.219:1081 | ✓ 491ms | ✓ 1109ms | ✓ 1159ms | ✓ 1355ms | ✓ 848ms | http |
| 185.125.100.115:40000 | ✓ 1222ms | ✓ 1950ms | 否 | ✓ 1694ms | ✓ 1856ms | http |
| 43.133.44.89:8888 | ✓ 1942ms | 否 | 否 | ✓ 1812ms | ✓ 1764ms | http |
| 45.116.14.87:8080 | 否 | 否 | ✓ 930ms | ✓ 1021ms | ✓ 742ms | http |
| 8.154.21.175:3128 | 否 | ✓ 1052ms | ✓ 1592ms | ✓ 1122ms | 否 | http |
| 121.230.9.19:1080 | 否 | ✓ 1335ms | ✓ 1037ms | ✓ 1290ms | 否 | http |
| 101.32.244.83:8080 | ✓ 956ms | 否 | ✓ 956ms | ✓ 1464ms | ✓ 1292ms | http |
| 121.43.196.213:8222 | ✓ 968ms | ✓ 1131ms | ✓ 828ms | ✓ 1174ms | ✓ 968ms | http |
| 121.43.196.210:8222 | ✓ 947ms | ✓ 1157ms | ✓ 858ms | ✓ 1142ms | ✓ 996ms | http |
| 101.32.243.189:80 | ✓ 1225ms | 否 | ✓ 1238ms | ✓ 1454ms | ✓ 1247ms | http |
| 103.157.200.126:3128 | ✓ 1187ms | 否 | ✓ 1208ms | 否 | ✓ 1556ms | http |
| 185.121.13.73:3128 | ✓ 1142ms | 否 | ✓ 1480ms | 否 | ✓ 1586ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1023ms | ✓ 311ms | ✓ 1055ms | 否 | http |
| 193.160.209.58:1080 | ✓ 1475ms | 否 | ✓ 1598ms | 否 | ✓ 1678ms | http |
| 64.188.77.26:3128 | ✓ 1133ms | 否 | ✓ 979ms | 否 | ✓ 1521ms | http |
| 223.16.170.103:80 | ✓ 1187ms | 否 | ✓ 1220ms | ✓ 1160ms | ✓ 1119ms | http |
| 3.101.133.120:80 | ✓ 904ms | ✓ 1375ms | ✓ 1596ms | ✓ 1291ms | ✓ 1104ms | http |
| 103.158.242.58:83 | ✓ 1697ms | 否 | ✓ 1657ms | ✓ 1843ms | 否 | http |
| 86.104.74.110:1082 | ✓ 1066ms | ✓ 1239ms | ✓ 964ms | ✓ 1729ms | ✓ 1312ms | http |
| 86.104.74.110:1081 | ✓ 1064ms | ✓ 1306ms | ✓ 898ms | ✓ 1728ms | ✓ 1330ms | http |
| 213.111.146.36:18080 | 否 | 否 | ✓ 1323ms | ✓ 1632ms | ✓ 1358ms | http |
| 116.171.106.111:3443 | 否 | 否 | ✓ 1463ms | ✓ 1884ms | ✓ 1606ms | http |
| 91.233.223.147:3128 | ✓ 1233ms | 否 | ✓ 1122ms | 否 | ✓ 1569ms | http |
| 80.92.204.47:1082 | ✓ 1659ms | 否 | ✓ 776ms | ✓ 1407ms | ✓ 1714ms | http |
| 2.27.32.81:3128 | ✓ 1168ms | 否 | ✓ 1002ms | 否 | ✓ 1678ms | http |
| 94.131.118.129:1081 | ✓ 1005ms | ✓ 1559ms | ✓ 1079ms | 否 | ✓ 1094ms | http |
| 38.194.254.134:999 | ✓ 1191ms | ✓ 1639ms | ✓ 1394ms | 否 | 否 | http |
| 107.174.64.143:1080 | ✓ 513ms | ✓ 1317ms | ✓ 991ms | ✓ 1206ms | ✓ 1098ms | http |
| 152.42.177.32:8888 | ✓ 981ms | 否 | ✓ 975ms | ✓ 1340ms | ✓ 1320ms | http |
| 121.230.8.253:1080 | ✓ 1509ms | ✓ 1719ms | 否 | ✓ 1567ms | 否 | http |
| 61.52.131.172:8443 | 否 | 否 | ✓ 1718ms | ✓ 1210ms | ✓ 1012ms | http |
| 103.172.70.173:8080 | ✓ 1819ms | 否 | ✓ 1442ms | 否 | ✓ 1382ms | http |
| 138.197.68.35:4857 | 否 | 否 | ✓ 660ms | ✓ 1329ms | ✓ 991ms | http |
| 103.157.117.116:8080 | ✓ 1798ms | 否 | 否 | ✓ 1768ms | ✓ 1714ms | http |
| 200.125.171.254:999 | 否 | ✓ 1692ms | ✓ 1242ms | ✓ 1798ms | ✓ 1557ms | http |

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
