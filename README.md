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

最后更新：2026-02-27 19:48:46 UTC（2026-02-28 03:48:46 UTC+8）

**代理总数：75**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 870ms | 否 | ✓ 852ms | ✓ 1599ms | ✓ 1025ms | http |
| 35.225.22.61:80 | 否 | ✓ 1899ms | ✓ 1200ms | ✓ 1354ms | ✓ 1127ms | http |
| 120.92.212.16:7890 | ✓ 1050ms | ✓ 1168ms | ✓ 1170ms | ✓ 1195ms | ✓ 931ms | http |
| 147.45.216.148:1080 | ✓ 677ms | 否 | ✓ 1450ms | 否 | ✓ 1649ms | http |
| 120.92.212.16:8890 | ✓ 1310ms | 否 | ✓ 1181ms | 否 | ✓ 914ms | http |
| 5.101.0.233:3128 | ✓ 1130ms | ✓ 1968ms | ✓ 1683ms | 否 | ✓ 1391ms | http |
| 103.84.95.54:7890 | ✓ 942ms | 否 | 否 | ✓ 771ms | ✓ 1127ms | http |
| 85.208.108.43:2094 | ✓ 391ms | 否 | ✓ 952ms | ✓ 1299ms | ✓ 859ms | http |
| 59.46.216.131:30001 | ✓ 1998ms | 否 | ✓ 1082ms | 否 | ✓ 993ms | http |
| 14.143.222.113:10158 | ✓ 963ms | 否 | ✓ 951ms | ✓ 1288ms | 否 | http |
| 158.178.237.243:3128 | ✓ 1125ms | 否 | ✓ 1901ms | ✓ 1479ms | ✓ 1036ms | http |
| 36.147.78.166:80 | 否 | ✓ 1558ms | ✓ 1609ms | ✓ 1647ms | 否 | http |
| 121.237.181.137:8888 | ✓ 1524ms | ✓ 1019ms | ✓ 792ms | 否 | ✓ 867ms | http |
| 101.43.255.96:80 | ✓ 1367ms | ✓ 1445ms | ✓ 1740ms | ✓ 1509ms | ✓ 988ms | http |
| 81.70.169.194:80 | ✓ 1253ms | ✓ 1632ms | ✓ 1940ms | ✓ 1606ms | ✓ 1231ms | http |
| 101.47.73.135:3128 | ✓ 1486ms | 否 | 否 | ✓ 1111ms | ✓ 1450ms | http |
| 103.215.36.88:19328 | ✓ 939ms | ✓ 1329ms | ✓ 1107ms | ✓ 1345ms | ✓ 1009ms | http |
| 47.110.42.192:9003 | ✓ 1402ms | ✓ 1337ms | ✓ 1286ms | ✓ 1463ms | ✓ 1263ms | http |
| 52.188.28.218:3128 | ✓ 1474ms | ✓ 1295ms | ✓ 1392ms | 否 | ✓ 882ms | http |
| 47.105.98.23:3128 | ✓ 1706ms | ✓ 1973ms | 否 | 否 | ✓ 1959ms | http |
| 170.78.208.245:999 | ✓ 1040ms | ✓ 1884ms | 否 | ✓ 1425ms | ✓ 1143ms | http |
| 35.234.17.221:8080 | 否 | 否 | ✓ 1170ms | ✓ 1410ms | ✓ 810ms | http |
| 14.56.107.244:3128 | ✓ 1585ms | 否 | 否 | ✓ 1021ms | ✓ 766ms | http |
| 61.72.110.24:3128 | ✓ 1813ms | 否 | ✓ 620ms | ✓ 1174ms | 否 | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1710ms | ✓ 1414ms | ✓ 1453ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1618ms | ✓ 1422ms | ✓ 1540ms | http |
| 168.235.110.63:3128 | ✓ 854ms | 否 | ✓ 1372ms | 否 | ✓ 1407ms | http |
| 210.223.44.230:3128 | 否 | ✓ 985ms | ✓ 980ms | ✓ 941ms | ✓ 1009ms | http |
| 34.101.184.164:3128 | ✓ 1413ms | 否 | ✓ 739ms | ✓ 1213ms | ✓ 952ms | http |
| 62.113.119.14:8080 | ✓ 1583ms | 否 | ✓ 1187ms | ✓ 1657ms | ✓ 1415ms | http |
| 91.238.104.172:2024 | ✓ 1356ms | 否 | 否 | ✓ 1949ms | ✓ 1433ms | http |
| 157.20.128.247:3125 | ✓ 1410ms | 否 | 否 | ✓ 1696ms | ✓ 1234ms | http |
| 121.230.9.198:1080 | ✓ 1132ms | ✓ 1453ms | ✓ 1206ms | ✓ 1266ms | ✓ 980ms | http |
| 165.227.5.10:8888 | ✓ 673ms | 否 | ✓ 369ms | 否 | ✓ 654ms | http |
| 36.147.78.166:443 | ✓ 1696ms | 否 | ✓ 1612ms | ✓ 1474ms | ✓ 1518ms | http |
| 121.230.9.184:1080 | ✓ 1432ms | ✓ 1476ms | 否 | ✓ 1797ms | 否 | http |
| 185.243.218.43:49153 | ✓ 1691ms | 否 | ✓ 1943ms | 否 | ✓ 1667ms | http |
| 121.230.8.153:1080 | ✓ 1214ms | ✓ 1595ms | ✓ 1049ms | ✓ 1522ms | ✓ 1285ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1512ms | ✓ 1648ms | ✓ 1334ms | ✓ 1342ms | http |
| 45.136.198.40:3128 | ✓ 1486ms | ✓ 1739ms | ✓ 1880ms | 否 | 否 | http |
| 175.215.60.67:3172 | 否 | ✓ 845ms | ✓ 811ms | ✓ 971ms | 否 | http |
| 175.215.60.67:3136 | 否 | ✓ 977ms | ✓ 925ms | ✓ 972ms | 否 | http |
| 175.215.60.137:3020 | 否 | ✓ 1210ms | ✓ 679ms | ✓ 972ms | 否 | http |
| 175.215.38.217:3068 | 否 | ✓ 801ms | ✓ 854ms | ✓ 976ms | 否 | http |
| 85.208.108.43:10808 | 否 | 否 | ✓ 1320ms | ✓ 1353ms | ✓ 998ms | http |
| 120.46.152.136:3128 | ✓ 994ms | ✓ 1335ms | ✓ 1386ms | ✓ 1641ms | ✓ 1075ms | http |
| 91.238.104.171:2023 | ✓ 844ms | ✓ 1757ms | ✓ 787ms | ✓ 1831ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1201ms | 否 | ✓ 1189ms | ✓ 1857ms | ✓ 897ms | http |
| 45.88.0.116:3128 | ✓ 1769ms | ✓ 1659ms | ✓ 1006ms | 否 | ✓ 1297ms | http |
| 45.88.0.117:3128 | ✓ 1702ms | ✓ 1596ms | ✓ 1826ms | ✓ 1589ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1743ms | 否 | ✓ 1456ms | 否 | ✓ 1204ms | http |
| 61.72.110.94:3128 | ✓ 1747ms | ✓ 1410ms | 否 | 否 | ✓ 1262ms | http |
| 103.155.166.150:8181 | ✓ 1664ms | 否 | ✓ 1477ms | ✓ 1319ms | ✓ 1225ms | http |
| 121.128.121.134:3128 | ✓ 1940ms | 否 | ✓ 1551ms | 否 | ✓ 790ms | http |
| 44.194.14.192:80 | ✓ 1835ms | ✓ 1713ms | ✓ 817ms | ✓ 1490ms | ✓ 958ms | http |
| 100.51.83.120:80 | ✓ 1014ms | ✓ 1733ms | ✓ 1488ms | ✓ 1451ms | ✓ 1336ms | http |
| 98.88.69.176:80 | 否 | 否 | ✓ 1586ms | ✓ 1500ms | ✓ 898ms | http |
| 103.82.23.118:5178 | ✓ 1988ms | 否 | 否 | ✓ 1821ms | ✓ 1338ms | http |
| 138.124.53.25:7443 | ✓ 1437ms | ✓ 1787ms | 否 | 否 | ✓ 1792ms | http |
| 103.236.64.247:8888 | 否 | ✓ 1552ms | ✓ 942ms | 否 | ✓ 1951ms | http |
| 14.56.118.164:3128 | 否 | ✓ 1791ms | 否 | ✓ 1294ms | ✓ 1912ms | http |
| 132.145.93.138:1080 | ✓ 1756ms | 否 | ✓ 1831ms | 否 | ✓ 1900ms | http |
| 8.219.97.248:80 | ✓ 1648ms | 否 | ✓ 1020ms | 否 | ✓ 1739ms | http |
| 121.230.8.136:1080 | 否 | 否 | ✓ 1563ms | ✓ 1707ms | ✓ 1260ms | http |
| 103.39.51.190:8080 | ✓ 1642ms | 否 | 否 | ✓ 1404ms | ✓ 1960ms | http |
| 14.56.118.24:3128 | ✓ 1946ms | 否 | ✓ 1988ms | 否 | ✓ 1348ms | http |
| 150.249.255.91:3128 | ✓ 1676ms | ✓ 919ms | ✓ 1123ms | ✓ 1207ms | ✓ 718ms | http |
| 61.72.110.54:3128 | ✓ 1700ms | 否 | 否 | ✓ 1159ms | ✓ 1949ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1064ms | ✓ 1461ms | ✓ 890ms | http |
| 121.128.121.34:3128 | 否 | 否 | ✓ 824ms | ✓ 1480ms | ✓ 796ms | http |
| 14.56.118.154:3128 | ✓ 1414ms | ✓ 1833ms | ✓ 1036ms | 否 | 否 | http |
| 45.140.147.82:1081 | ✓ 1573ms | 否 | ✓ 1304ms | 否 | ✓ 1431ms | http |
| 45.186.6.104:3128 | ✓ 1185ms | ✓ 1997ms | ✓ 1800ms | 否 | 否 | http |
| 103.82.23.118:5247 | ✓ 1814ms | 否 | ✓ 1268ms | ✓ 1774ms | ✓ 1752ms | http |
| 112.65.132.182:3128 | ✓ 1023ms | 否 | ✓ 1809ms | ✓ 1016ms | ✓ 810ms | http |

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
