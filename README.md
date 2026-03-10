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

最后更新：2026-03-10 11:39:24 UTC（2026-03-10 19:39:24 UTC+8）

**代理总数：67**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | 否 | ✓ 1769ms | ✓ 1632ms | ✓ 1057ms | ✓ 790ms | http |
| 154.3.236.202:3128 | 否 | 否 | ✓ 1407ms | ✓ 1342ms | ✓ 865ms | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1934ms | ✓ 1457ms | ✓ 1474ms | http |
| 217.76.245.80:999 | ✓ 761ms | 否 | ✓ 1514ms | ✓ 1270ms | ✓ 1398ms | http |
| 190.212.131.238:3128 | ✓ 1595ms | ✓ 1879ms | ✓ 982ms | 否 | ✓ 1981ms | http |
| 91.107.141.42:8081 | ✓ 1732ms | ✓ 1502ms | ✓ 1249ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1634ms | ✓ 1505ms | 否 | ✓ 1718ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1149ms | 否 | ✓ 1926ms | ✓ 1767ms | 否 | http |
| 81.70.169.194:80 | 否 | ✓ 1458ms | ✓ 1079ms | 否 | ✓ 1260ms | http |
| 190.9.109.198:999 | ✓ 911ms | ✓ 1379ms | ✓ 996ms | ✓ 1353ms | ✓ 1235ms | http |
| 202.155.12.161:443 | ✓ 1845ms | 否 | ✓ 1338ms | ✓ 1354ms | ✓ 1164ms | http |
| 152.42.213.210:8080 | ✓ 1334ms | 否 | ✓ 1389ms | ✓ 1342ms | ✓ 1005ms | http |
| 101.43.255.96:80 | 否 | 否 | ✓ 1483ms | ✓ 1782ms | ✓ 1137ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1411ms | ✓ 1109ms | ✓ 1440ms | 否 | http |
| 194.213.18.200:443 | ✓ 1621ms | 否 | ✓ 1173ms | 否 | ✓ 1823ms | http |
| 101.47.73.135:3128 | ✓ 1532ms | 否 | ✓ 1986ms | ✓ 1581ms | ✓ 1323ms | http |
| 117.159.239.54:22222 | ✓ 1246ms | ✓ 1273ms | ✓ 1104ms | ✓ 1329ms | 否 | http |
| 45.140.147.155:1082 | ✓ 521ms | 否 | ✓ 1965ms | ✓ 1514ms | ✓ 1395ms | http |
| 120.92.212.16:8890 | ✓ 1135ms | ✓ 1452ms | ✓ 1184ms | ✓ 1766ms | 否 | http |
| 114.55.226.123:10086 | ✓ 1963ms | ✓ 1646ms | ✓ 1148ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1863ms | ✓ 1623ms | ✓ 1680ms | ✓ 1366ms | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1275ms | ✓ 415ms | ✓ 987ms | 否 | http |
| 45.136.130.223:8443 | ✓ 411ms | ✓ 978ms | ✓ 1063ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 796ms | ✓ 1248ms | ✓ 882ms | ✓ 1560ms | ✓ 1054ms | http |
| 168.235.110.63:3128 | ✓ 438ms | ✓ 939ms | ✓ 957ms | ✓ 1023ms | ✓ 793ms | http |
| 95.3.9.78:3128 | ✓ 1255ms | 否 | ✓ 1956ms | ✓ 1780ms | 否 | http |
| 165.227.5.10:8888 | ✓ 957ms | 否 | ✓ 1327ms | 否 | ✓ 975ms | http |
| 59.46.216.131:30001 | ✓ 1980ms | ✓ 1544ms | 否 | ✓ 1748ms | 否 | http |
| 95.3.9.78:8080 | ✓ 765ms | 否 | ✓ 628ms | ✓ 1556ms | ✓ 1214ms | http |
| 101.32.244.83:8080 | ✓ 1201ms | ✓ 1602ms | ✓ 1155ms | ✓ 1516ms | ✓ 1259ms | http |
| 46.183.25.8:443 | ✓ 1383ms | 否 | 否 | ✓ 1900ms | ✓ 1772ms | http |
| 185.191.236.162:3128 | ✓ 621ms | ✓ 1763ms | ✓ 1090ms | 否 | 否 | http |
| 152.42.213.210:80 | ✓ 1802ms | 否 | ✓ 1006ms | ✓ 1629ms | ✓ 1251ms | http |
| 45.136.198.40:3128 | ✓ 1721ms | ✓ 1814ms | 否 | ✓ 1973ms | ✓ 1548ms | http |
| 47.77.193.180:1080 | ✓ 929ms | ✓ 1160ms | ✓ 800ms | ✓ 1688ms | 否 | http |
| 45.136.131.47:8443 | ✓ 331ms | ✓ 1471ms | ✓ 1037ms | ✓ 1142ms | ✓ 857ms | http |
| 121.138.61.193:8933 | ✓ 997ms | ✓ 1411ms | ✓ 1180ms | ✓ 1246ms | ✓ 1289ms | http |
| 152.70.98.46:8888 | ✓ 1480ms | 否 | ✓ 1846ms | ✓ 1206ms | ✓ 1047ms | http |
| 103.173.139.2:8080 | 否 | 否 | ✓ 1912ms | ✓ 1874ms | ✓ 1703ms | http |
| 45.140.147.82:1081 | 否 | 否 | ✓ 840ms | ✓ 1780ms | ✓ 1113ms | http |
| 45.22.209.157:8888 | ✓ 1601ms | 否 | 否 | ✓ 1477ms | ✓ 1030ms | http |
| 45.186.6.104:3128 | ✓ 1675ms | ✓ 1860ms | ✓ 1942ms | 否 | 否 | http |
| 35.206.88.200:8888 | ✓ 794ms | 否 | ✓ 1396ms | ✓ 1155ms | 否 | http |
| 38.180.2.107:3128 | ✓ 1835ms | 否 | ✓ 1774ms | 否 | ✓ 1811ms | http |
| 45.140.147.82:1082 | ✓ 865ms | 否 | ✓ 913ms | ✓ 1219ms | ✓ 1164ms | http |
| 129.226.155.60:3128 | ✓ 1809ms | 否 | ✓ 1316ms | ✓ 1329ms | ✓ 1047ms | http |
| 14.225.222.164:7890 | 否 | 否 | ✓ 1637ms | ✓ 1363ms | ✓ 1156ms | http |
| 1.231.81.166:3128 | ✓ 1859ms | ✓ 1848ms | ✓ 1730ms | ✓ 1415ms | ✓ 1271ms | http |
| 202.129.206.239:3128 | ✓ 1920ms | 否 | ✓ 1948ms | ✓ 1780ms | ✓ 1786ms | http |
| 120.240.35.176:22222 | ✓ 1093ms | ✓ 1419ms | ✓ 1424ms | ✓ 1367ms | ✓ 1125ms | http |
| 183.249.5.111:22222 | 否 | ✓ 1415ms | ✓ 897ms | 否 | ✓ 942ms | http |
| 183.249.5.109:22222 | ✓ 978ms | ✓ 1379ms | ✓ 866ms | ✓ 1283ms | ✓ 1154ms | http |
| 120.240.35.177:22222 | ✓ 1515ms | 否 | ✓ 1546ms | ✓ 1544ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1985ms | 否 | 否 | ✓ 1659ms | ✓ 1892ms | http |
| 186.116.148.52:8080 | ✓ 1342ms | 否 | ✓ 1889ms | 否 | ✓ 1373ms | http |
| 117.159.239.58:22222 | ✓ 1212ms | ✓ 1279ms | ✓ 1218ms | ✓ 1328ms | ✓ 1150ms | http |
| 120.240.35.178:22222 | ✓ 1300ms | ✓ 1527ms | ✓ 1460ms | ✓ 1565ms | ✓ 1585ms | http |
| 159.223.42.219:3128 | ✓ 1524ms | 否 | ✓ 971ms | ✓ 1588ms | ✓ 1516ms | http |
| 120.238.159.229:22222 | ✓ 1090ms | ✓ 1456ms | ✓ 1151ms | ✓ 1354ms | ✓ 1125ms | http |
| 117.159.239.49:22222 | ✓ 1025ms | ✓ 1340ms | ✓ 1024ms | ✓ 1669ms | 否 | http |
| 113.59.32.161:22222 | 否 | ✓ 1490ms | ✓ 1273ms | ✓ 1368ms | ✓ 1170ms | http |
| 121.43.196.210:8222 | ✓ 1205ms | ✓ 1262ms | ✓ 1056ms | ✓ 1345ms | ✓ 1019ms | http |
| 121.43.196.213:8222 | ✓ 1149ms | ✓ 1261ms | ✓ 1093ms | ✓ 1309ms | ✓ 1106ms | http |
| 61.52.131.172:8443 | ✓ 1057ms | ✓ 1366ms | ✓ 1113ms | ✓ 1366ms | ✓ 1112ms | http |
| 120.238.159.230:22222 | ✓ 1096ms | ✓ 1463ms | 否 | ✓ 1337ms | ✓ 1064ms | http |
| 185.41.152.110:3128 | ✓ 1210ms | 否 | 否 | ✓ 1802ms | ✓ 1307ms | http |
| 37.139.33.145:1080 | ✓ 1715ms | 否 | ✓ 1515ms | 否 | ✓ 1791ms | http |

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
