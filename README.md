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

最后更新：2026-03-30 10:09:36 UTC（2026-03-30 18:09:36 UTC+8）

**代理总数：69**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 885ms | ✓ 1192ms | ✓ 827ms | ✓ 1056ms | ✓ 871ms | http |
| 39.185.46.193:5911 | ✓ 838ms | ✓ 1298ms | ✓ 1091ms | ✓ 1112ms | ✓ 1041ms | http |
| 147.161.239.240:8800 | ✓ 1009ms | ✓ 1493ms | ✓ 1266ms | ✓ 1858ms | ✓ 1427ms | http |
| 95.213.217.168:52004 | ✓ 1046ms | ✓ 1857ms | ✓ 1212ms | 否 | ✓ 1604ms | http |
| 147.161.210.140:8800 | ✓ 1845ms | 否 | ✓ 1049ms | ✓ 1749ms | ✓ 1169ms | http |
| 1.231.81.166:3128 | ✓ 1879ms | ✓ 1450ms | ✓ 1799ms | ✓ 1409ms | ✓ 1084ms | http |
| 167.103.115.102:8800 | ✓ 1155ms | 否 | ✓ 1241ms | ✓ 1659ms | ✓ 1545ms | http |
| 167.103.34.108:8800 | 否 | 否 | ✓ 1377ms | ✓ 1624ms | ✓ 1844ms | http |
| 222.184.48.242:22222 | ✓ 1162ms | ✓ 1531ms | 否 | ✓ 1405ms | ✓ 1280ms | http |
| 167.103.144.127:8800 | ✓ 1147ms | 否 | ✓ 1323ms | ✓ 1639ms | ✓ 1494ms | http |
| 167.103.31.122:8800 | ✓ 1252ms | 否 | ✓ 1571ms | ✓ 1742ms | ✓ 1639ms | http |
| 115.231.181.40:8128 | ✓ 1114ms | ✓ 1886ms | ✓ 1069ms | 否 | 否 | http |
| 208.87.243.199:7878 | ✓ 404ms | ✓ 1435ms | 否 | ✓ 1986ms | 否 | http |
| 198.59.68.130:3128 | 否 | 否 | ✓ 1678ms | ✓ 1154ms | ✓ 1487ms | http |
| 85.208.108.43:2094 | ✓ 461ms | 否 | ✓ 459ms | ✓ 1015ms | ✓ 843ms | http |
| 222.184.48.251:22222 | ✓ 1078ms | ✓ 1472ms | ✓ 1166ms | ✓ 1481ms | ✓ 1332ms | http |
| 222.184.48.241:22222 | ✓ 1608ms | ✓ 1347ms | 否 | ✓ 1412ms | ✓ 1054ms | http |
| 101.43.127.100:8877 | ✓ 1079ms | ✓ 1410ms | ✓ 1070ms | 否 | ✓ 1134ms | http |
| 177.234.217.88:999 | 否 | 否 | ✓ 1715ms | ✓ 1776ms | ✓ 1450ms | http |
| 45.149.92.147:5001 | ✓ 826ms | 否 | 否 | ✓ 1139ms | ✓ 1037ms | http |
| 113.160.132.26:8080 | ✓ 1755ms | 否 | ✓ 1497ms | ✓ 1517ms | ✓ 1585ms | http |
| 45.88.0.111:3128 | 否 | ✓ 1543ms | 否 | ✓ 1545ms | ✓ 1288ms | http |
| 45.88.0.117:3128 | 否 | 否 | ✓ 1819ms | ✓ 1262ms | ✓ 1290ms | http |
| 183.249.5.105:22222 | 否 | 否 | ✓ 1197ms | ✓ 1395ms | ✓ 1057ms | http |
| 183.249.5.117:22222 | 否 | 否 | ✓ 1092ms | ✓ 1483ms | ✓ 1131ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 660ms | ✓ 1005ms | ✓ 1808ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1332ms | ✓ 1601ms | ✓ 1243ms | http |
| 45.12.151.226:2829 | ✓ 965ms | ✓ 1624ms | ✓ 859ms | 否 | ✓ 1208ms | http |
| 35.225.22.61:80 | ✓ 346ms | ✓ 1333ms | ✓ 766ms | ✓ 1004ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1197ms | ✓ 1428ms | ✓ 1105ms | 否 | 否 | http |
| 209.126.84.232:8888 | ✓ 1465ms | ✓ 1686ms | ✓ 1268ms | ✓ 1429ms | ✓ 1168ms | http |
| 101.47.73.135:3128 | ✓ 1677ms | 否 | ✓ 1886ms | ✓ 1596ms | 否 | http |
| 46.39.105.157:8080 | ✓ 1147ms | 否 | ✓ 1201ms | 否 | ✓ 1536ms | http |
| 103.84.95.54:7890 | ✓ 1442ms | 否 | ✓ 983ms | 否 | ✓ 992ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1435ms | ✓ 1171ms | ✓ 1413ms | 否 | http |
| 101.255.211.106:3128 | ✓ 1574ms | 否 | ✓ 1970ms | 否 | ✓ 1721ms | http |
| 101.32.244.83:8080 | ✓ 1164ms | 否 | ✓ 1187ms | ✓ 1536ms | ✓ 1490ms | http |
| 103.166.185.54:3128 | 否 | ✓ 1781ms | ✓ 1699ms | ✓ 1443ms | ✓ 1210ms | http |
| 183.249.5.110:22222 | ✓ 1054ms | ✓ 1334ms | ✓ 851ms | ✓ 1149ms | ✓ 899ms | http |
| 141.144.231.186:3128 | ✓ 839ms | 否 | ✓ 1737ms | ✓ 1849ms | ✓ 1544ms | http |
| 222.184.48.252:22222 | ✓ 1143ms | ✓ 1297ms | ✓ 1730ms | ✓ 1438ms | ✓ 1046ms | http |
| 5.102.109.41:999 | ✓ 781ms | ✓ 1272ms | 否 | ✓ 1438ms | ✓ 1143ms | http |
| 42.96.16.158:1311 | ✓ 1350ms | 否 | ✓ 1520ms | ✓ 1492ms | ✓ 1835ms | http |
| 31.192.106.135:8010 | ✓ 1267ms | 否 | ✓ 1757ms | 否 | ✓ 1512ms | http |
| 116.80.49.166:3172 | ✓ 1796ms | 否 | ✓ 1712ms | 否 | ✓ 1870ms | http |
| 194.59.204.87:9080 | ✓ 539ms | ✓ 1623ms | ✓ 1129ms | 否 | 否 | http |
| 45.88.0.115:3128 | 否 | 否 | ✓ 1150ms | ✓ 1552ms | ✓ 1281ms | http |
| 213.220.62.62:3128 | 否 | 否 | ✓ 507ms | ✓ 1265ms | ✓ 966ms | http |
| 210.223.44.230:3128 | ✓ 828ms | ✓ 1444ms | ✓ 941ms | 否 | ✓ 948ms | http |
| 62.171.161.88:2018 | ✓ 793ms | 否 | ✓ 663ms | ✓ 1471ms | ✓ 1110ms | http |
| 94.158.49.82:3128 | ✓ 1443ms | 否 | 否 | ✓ 1880ms | ✓ 1951ms | http |
| 45.88.0.99:3128 | ✓ 1737ms | 否 | ✓ 1924ms | ✓ 1974ms | ✓ 1369ms | http |
| 45.88.0.116:3128 | ✓ 931ms | 否 | ✓ 707ms | ✓ 1305ms | 否 | http |
| 217.76.245.80:999 | ✓ 606ms | ✓ 1505ms | 否 | 否 | ✓ 1005ms | http |
| 181.78.44.63:999 | ✓ 1004ms | ✓ 1622ms | ✓ 1496ms | 否 | ✓ 1947ms | http |
| 121.43.196.213:8222 | ✓ 1182ms | ✓ 1284ms | ✓ 1034ms | ✓ 1346ms | ✓ 1062ms | http |
| 121.43.196.210:8222 | ✓ 1244ms | ✓ 1288ms | ✓ 1040ms | ✓ 1344ms | ✓ 1059ms | http |
| 114.55.226.123:10086 | ✓ 1925ms | ✓ 1606ms | ✓ 1251ms | ✓ 1505ms | ✓ 1216ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1628ms | ✓ 1777ms | ✓ 1600ms | http |
| 38.180.2.107:3128 | ✓ 724ms | ✓ 1593ms | ✓ 1532ms | 否 | 否 | http |
| 160.238.65.9:3128 | 否 | 否 | ✓ 484ms | ✓ 1575ms | ✓ 983ms | http |
| 160.238.65.3:3128 | ✓ 1833ms | 否 | ✓ 423ms | ✓ 1242ms | 否 | http |
| 160.238.65.5:3128 | ✓ 891ms | 否 | ✓ 1593ms | ✓ 1256ms | 否 | http |
| 160.238.65.7:3128 | ✓ 894ms | 否 | ✓ 1361ms | ✓ 1246ms | ✓ 1023ms | http |
| 160.238.65.6:3128 | ✓ 892ms | 否 | ✓ 1593ms | ✓ 1872ms | 否 | http |
| 59.8.203.55:80 | ✓ 1712ms | ✓ 1390ms | ✓ 1211ms | ✓ 1268ms | ✓ 963ms | http |
| 62.113.119.14:8080 | ✓ 901ms | ✓ 1813ms | ✓ 990ms | 否 | 否 | http |
| 183.249.5.111:22222 | ✓ 933ms | ✓ 1353ms | ✓ 850ms | ✓ 1133ms | ✓ 900ms | http |
| 180.130.80.196:9003 | ✓ 1432ms | ✓ 1654ms | ✓ 1595ms | ✓ 1560ms | ✓ 1793ms | http |

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
