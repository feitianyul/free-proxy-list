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

最后更新：2026-03-01 00:25:32 UTC（2026-03-01 08:25:32 UTC+8）

**代理总数：89**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | 否 | ✓ 1193ms | ✓ 679ms | ✓ 1194ms | ✓ 1049ms | http |
| 148.135.85.87:1080 | ✓ 1089ms | 否 | ✓ 1099ms | ✓ 1071ms | ✓ 836ms | http |
| 34.78.200.22:3128 | ✓ 993ms | 否 | ✓ 1237ms | ✓ 1768ms | ✓ 1242ms | http |
| 34.32.154.33:3128 | ✓ 974ms | ✓ 1673ms | ✓ 1262ms | 否 | ✓ 1742ms | http |
| 5.101.0.233:3128 | ✓ 1004ms | 否 | ✓ 1355ms | ✓ 1922ms | ✓ 1237ms | http |
| 3.213.157.4:3128 | ✓ 292ms | 否 | ✓ 1237ms | ✓ 1212ms | ✓ 1393ms | http |
| 34.79.102.160:3128 | ✓ 484ms | ✓ 1650ms | ✓ 1089ms | ✓ 1742ms | ✓ 1594ms | http |
| 34.78.177.18:3128 | ✓ 630ms | ✓ 1772ms | ✓ 1002ms | 否 | ✓ 1696ms | http |
| 34.89.174.168:3128 | ✓ 561ms | 否 | ✓ 1507ms | ✓ 1996ms | ✓ 1614ms | http |
| 34.159.121.205:3128 | ✓ 578ms | ✓ 1882ms | ✓ 1550ms | 否 | ✓ 1702ms | http |
| 34.7.88.87:3128 | ✓ 638ms | ✓ 1952ms | ✓ 1547ms | 否 | ✓ 1889ms | http |
| 34.158.73.60:3128 | 否 | 否 | ✓ 1404ms | ✓ 1851ms | ✓ 1397ms | http |
| 36.147.78.166:80 | 否 | ✓ 1842ms | ✓ 1796ms | ✓ 1909ms | 否 | http |
| 104.238.30.91:63900 | ✓ 1953ms | 否 | ✓ 1968ms | 否 | ✓ 1995ms | http |
| 34.185.159.217:3128 | ✓ 707ms | ✓ 1894ms | ✓ 1497ms | 否 | 否 | http |
| 168.235.110.63:3128 | ✓ 1905ms | 否 | ✓ 988ms | ✓ 1063ms | ✓ 820ms | http |
| 52.188.28.218:3128 | 否 | ✓ 1918ms | ✓ 125ms | ✓ 885ms | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1691ms | ✓ 900ms | ✓ 983ms | ✓ 793ms | http |
| 103.84.95.54:7890 | ✓ 897ms | 否 | ✓ 752ms | 否 | ✓ 823ms | http |
| 103.215.36.88:13763 | ✓ 1188ms | ✓ 1393ms | ✓ 1171ms | 否 | 否 | http |
| 14.56.107.244:3128 | 否 | ✓ 1883ms | 否 | ✓ 1459ms | ✓ 1017ms | http |
| 39.104.201.40:7890 | ✓ 986ms | ✓ 1318ms | ✓ 1104ms | ✓ 1432ms | ✓ 1051ms | http |
| 101.43.255.96:80 | ✓ 1062ms | ✓ 1461ms | ✓ 1183ms | ✓ 1347ms | ✓ 1158ms | http |
| 81.70.169.194:80 | ✓ 1168ms | ✓ 1478ms | ✓ 1162ms | ✓ 1421ms | ✓ 1176ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 986ms | ✓ 1489ms | ✓ 1316ms | http |
| 120.92.212.16:7890 | ✓ 1817ms | 否 | ✓ 1212ms | 否 | ✓ 1097ms | http |
| 120.92.212.16:8890 | ✓ 1828ms | ✓ 1356ms | 否 | ✓ 1657ms | 否 | http |
| 121.237.181.137:8888 | ✓ 995ms | ✓ 1283ms | ✓ 1019ms | ✓ 1299ms | ✓ 908ms | http |
| 210.223.44.230:3128 | ✓ 1146ms | ✓ 1025ms | ✓ 1250ms | ✓ 1096ms | ✓ 1459ms | http |
| 111.79.111.126:3128 | ✓ 1340ms | 否 | 否 | ✓ 1966ms | ✓ 1712ms | http |
| 35.234.17.221:8080 | ✓ 1427ms | ✓ 1739ms | ✓ 1537ms | ✓ 1176ms | 否 | http |
| 143.198.37.6:8888 | 否 | ✓ 1093ms | ✓ 166ms | ✓ 917ms | ✓ 696ms | http |
| 74.208.234.198:443 | 否 | ✓ 1830ms | ✓ 1764ms | ✓ 1548ms | ✓ 1314ms | http |
| 103.86.131.62:80 | ✓ 1710ms | 否 | 否 | ✓ 1456ms | ✓ 1162ms | http |
| 38.180.2.107:3128 | ✓ 939ms | ✓ 1925ms | ✓ 1626ms | 否 | 否 | http |
| 185.115.74.185:8080 | ✓ 902ms | ✓ 1957ms | ✓ 1691ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 1084ms | ✓ 1591ms | ✓ 982ms | ✓ 1457ms | ✓ 1104ms | http |
| 59.46.216.131:30001 | ✓ 1085ms | ✓ 1494ms | ✓ 1292ms | ✓ 1512ms | ✓ 1474ms | http |
| 43.161.214.161:1080 | ✓ 1088ms | ✓ 1307ms | ✓ 1192ms | ✓ 1142ms | ✓ 1016ms | http |
| 104.238.30.38:59741 | ✓ 1693ms | 否 | ✓ 1867ms | 否 | ✓ 1999ms | http |
| 113.255.59.226:8080 | 否 | 否 | ✓ 1432ms | ✓ 1286ms | ✓ 1327ms | http |
| 61.72.110.94:3128 | ✓ 1366ms | 否 | 否 | ✓ 1437ms | ✓ 1283ms | http |
| 104.238.30.63:63744 | ✓ 1772ms | 否 | ✓ 1839ms | 否 | ✓ 1999ms | http |
| 91.238.104.171:2023 | ✓ 1185ms | ✓ 1547ms | ✓ 1565ms | 否 | 否 | http |
| 24.199.124.151:3128 | ✓ 496ms | ✓ 1677ms | ✓ 1146ms | ✓ 1012ms | ✓ 724ms | http |
| 104.238.30.40:59741 | ✓ 1752ms | 否 | ✓ 1839ms | 否 | ✓ 1999ms | http |
| 91.238.104.172:2024 | ✓ 1655ms | ✓ 1662ms | ✓ 1653ms | ✓ 1474ms | ✓ 1197ms | http |
| 162.240.154.26:3128 | ✓ 600ms | 否 | ✓ 922ms | ✓ 1086ms | ✓ 714ms | http |
| 138.124.53.25:7443 | ✓ 774ms | 否 | ✓ 1579ms | ✓ 1883ms | ✓ 1423ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1397ms | 否 | ✓ 1116ms | ✓ 985ms | http |
| 121.230.8.144:1080 | ✓ 1253ms | ✓ 1784ms | ✓ 1395ms | ✓ 1778ms | ✓ 1292ms | http |
| 49.7.179.70:3333 | 否 | ✓ 1628ms | ✓ 1041ms | ✓ 1480ms | ✓ 1105ms | http |
| 103.104.99.29:80 | ✓ 1921ms | 否 | ✓ 1818ms | ✓ 1674ms | ✓ 1636ms | http |
| 103.104.99.89:80 | ✓ 1922ms | 否 | ✓ 1794ms | ✓ 1656ms | ✓ 1825ms | http |
| 125.64.244.100:8889 | ✓ 1749ms | ✓ 1865ms | ✓ 1768ms | 否 | ✓ 1838ms | http |
| 8.217.147.173:8080 | 否 | 否 | ✓ 1371ms | ✓ 1834ms | ✓ 1372ms | http |
| 104.238.30.37:59741 | ✓ 1848ms | 否 | ✓ 1980ms | 否 | ✓ 1999ms | http |
| 103.35.188.243:3128 | 否 | ✓ 988ms | 否 | ✓ 1095ms | ✓ 820ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1463ms | ✓ 1477ms | ✓ 1275ms | http |
| 116.58.162.45:3128 | ✓ 1357ms | ✓ 1858ms | ✓ 1358ms | 否 | 否 | http |
| 180.127.149.244:1080 | ✓ 941ms | ✓ 1246ms | 否 | ✓ 1283ms | ✓ 1061ms | http |
| 61.72.110.54:3128 | ✓ 1767ms | 否 | ✓ 776ms | 否 | ✓ 936ms | http |
| 104.238.30.45:59741 | ✓ 1720ms | 否 | ✓ 1812ms | 否 | ✓ 1995ms | http |
| 104.238.30.50:59741 | ✓ 1720ms | 否 | ✓ 1807ms | 否 | ✓ 1999ms | http |
| 36.147.78.166:443 | ✓ 1864ms | 否 | ✓ 1727ms | 否 | ✓ 1798ms | http |
| 5.75.201.136:1080 | ✓ 970ms | ✓ 1977ms | ✓ 1549ms | ✓ 1968ms | ✓ 1798ms | http |
| 45.140.147.82:1082 | ✓ 596ms | ✓ 1876ms | 否 | ✓ 1222ms | 否 | http |
| 45.140.147.82:1081 | ✓ 597ms | ✓ 1490ms | ✓ 1482ms | ✓ 1514ms | ✓ 1123ms | http |
| 61.72.110.24:3128 | ✓ 744ms | 否 | ✓ 902ms | ✓ 1141ms | 否 | http |
| 103.236.64.247:8888 | 否 | ✓ 1397ms | ✓ 1053ms | ✓ 1288ms | 否 | http |
| 222.102.86.137:3028 | 否 | ✓ 1497ms | ✓ 1249ms | ✓ 1962ms | ✓ 1087ms | http |
| 222.102.86.137:3016 | 否 | ✓ 1513ms | ✓ 1950ms | 否 | ✓ 1324ms | http |
| 103.93.93.66:3125 | ✓ 1414ms | 否 | ✓ 1706ms | ✓ 1553ms | ✓ 1606ms | http |
| 14.143.222.113:10158 | ✓ 1035ms | 否 | ✓ 1092ms | ✓ 1445ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1865ms | 否 | 否 | ✓ 1532ms | ✓ 1596ms | http |
| 2.56.178.131:443 | ✓ 1448ms | 否 | ✓ 1605ms | 否 | ✓ 1945ms | http |
| 190.248.145.114:999 | ✓ 1095ms | 否 | 否 | ✓ 1768ms | ✓ 1961ms | http |
| 45.136.198.40:3128 | ✓ 725ms | ✓ 1516ms | ✓ 1645ms | 否 | ✓ 1732ms | http |
| 101.47.73.135:3128 | ✓ 1565ms | 否 | ✓ 1578ms | 否 | ✓ 1353ms | http |
| 51.79.71.106:8080 | ✓ 639ms | 否 | ✓ 1810ms | ✓ 1386ms | ✓ 1261ms | http |
| 104.238.30.39:59741 | ✓ 1766ms | 否 | ✓ 1807ms | 否 | ✓ 1967ms | http |
| 90.84.188.97:8000 | ✓ 968ms | ✓ 1703ms | ✓ 808ms | 否 | ✓ 1651ms | http |
| 192.71.213.85:9091 | ✓ 705ms | 否 | ✓ 665ms | ✓ 1643ms | 否 | http |
| 103.215.36.88:16316 | ✓ 1175ms | ✓ 1416ms | ✓ 1092ms | ✓ 1344ms | ✓ 1156ms | http |
| 34.101.184.164:3128 | ✓ 1744ms | 否 | ✓ 1801ms | ✓ 1719ms | ✓ 1622ms | http |
| 172.212.68.37:3128 | ✓ 319ms | ✓ 1658ms | ✓ 1227ms | ✓ 1357ms | ✓ 1469ms | http |
| 1.225.116.115:1080 | 否 | ✓ 1751ms | ✓ 1624ms | ✓ 1731ms | 否 | http |
| 103.131.19.42:8181 | ✓ 1662ms | 否 | ✓ 1488ms | 否 | ✓ 1560ms | http |
| 31.59.129.75:8080 | ✓ 886ms | ✓ 1535ms | 否 | 否 | ✓ 924ms | http |

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
