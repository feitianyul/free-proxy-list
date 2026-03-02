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

最后更新：2026-03-02 11:41:58 UTC（2026-03-02 19:41:58 UTC+8）

**代理总数：63**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 182ms | 否 | ✓ 837ms | ✓ 1073ms | ✓ 808ms | http |
| 45.140.147.82:1081 | ✓ 383ms | 否 | 否 | ✓ 1387ms | ✓ 1048ms | http |
| 5.75.196.26:40000 | ✓ 458ms | 否 | ✓ 1299ms | 否 | ✓ 1938ms | http |
| 91.238.104.171:2023 | ✓ 986ms | 否 | ✓ 1348ms | 否 | ✓ 1671ms | http |
| 121.128.121.54:3128 | 否 | 否 | ✓ 1117ms | ✓ 1267ms | ✓ 996ms | http |
| 59.46.216.131:30001 | ✓ 1996ms | ✓ 1455ms | 否 | 否 | ✓ 1524ms | http |
| 120.92.212.16:8890 | ✓ 1117ms | ✓ 1271ms | ✓ 1618ms | ✓ 1575ms | ✓ 1736ms | http |
| 91.238.104.172:2024 | ✓ 614ms | 否 | ✓ 951ms | ✓ 1456ms | ✓ 1397ms | http |
| 5.101.0.233:3128 | ✓ 861ms | 否 | ✓ 1097ms | ✓ 1919ms | ✓ 1389ms | http |
| 38.207.165.2:6005 | ✓ 1765ms | ✓ 1926ms | ✓ 1190ms | ✓ 1401ms | ✓ 863ms | http |
| 125.128.12.84:3128 | 否 | ✓ 1708ms | ✓ 1755ms | ✓ 1318ms | 否 | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1280ms | ✓ 1690ms | ✓ 1431ms | http |
| 61.72.110.54:3128 | ✓ 859ms | 否 | ✓ 1140ms | ✓ 1295ms | 否 | http |
| 14.56.177.44:3128 | ✓ 1431ms | ✓ 1745ms | ✓ 1094ms | ✓ 1266ms | ✓ 962ms | http |
| 101.43.255.96:80 | ✓ 1055ms | ✓ 1465ms | ✓ 1032ms | ✓ 1255ms | 否 | http |
| 194.87.43.49:8888 | ✓ 955ms | 否 | ✓ 1639ms | 否 | ✓ 1961ms | http |
| 81.70.169.194:80 | ✓ 1673ms | ✓ 1423ms | ✓ 1978ms | ✓ 1622ms | ✓ 1116ms | http |
| 62.113.119.14:8080 | ✓ 595ms | ✓ 1337ms | ✓ 901ms | ✓ 1395ms | ✓ 1087ms | http |
| 14.56.107.244:3128 | ✓ 1076ms | ✓ 1557ms | ✓ 1221ms | ✓ 1888ms | ✓ 1747ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1526ms | 否 | ✓ 1292ms | ✓ 1019ms | http |
| 61.72.221.94:3128 | 否 | ✓ 1474ms | ✓ 1223ms | ✓ 1579ms | ✓ 1002ms | http |
| 107.174.133.10:3128 | ✓ 665ms | ✓ 1522ms | ✓ 988ms | ✓ 1079ms | 否 | http |
| 195.123.209.48:3128 | ✓ 559ms | 否 | ✓ 485ms | ✓ 1415ms | ✓ 1159ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1684ms | 否 | ✓ 1073ms | ✓ 983ms | http |
| 101.47.73.135:3128 | ✓ 1623ms | 否 | ✓ 1746ms | ✓ 1581ms | ✓ 1370ms | http |
| 103.215.36.88:17100 | ✓ 965ms | ✓ 1855ms | ✓ 1118ms | 否 | 否 | http |
| 74.208.234.198:443 | ✓ 867ms | ✓ 1771ms | ✓ 1961ms | 否 | 否 | http |
| 45.140.147.82:1082 | ✓ 973ms | ✓ 1717ms | ✓ 1171ms | ✓ 1674ms | ✓ 1033ms | http |
| 125.128.12.14:3128 | ✓ 1276ms | 否 | ✓ 1055ms | ✓ 1608ms | ✓ 1643ms | http |
| 35.234.17.221:8080 | 否 | 否 | ✓ 1564ms | ✓ 1211ms | ✓ 1020ms | http |
| 171.234.62.116:10008 | ✓ 1963ms | 否 | 否 | ✓ 1986ms | ✓ 1476ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 997ms | ✓ 1138ms | ✓ 1086ms | http |
| 103.84.95.54:7890 | ✓ 816ms | 否 | ✓ 855ms | ✓ 1301ms | 否 | http |
| 61.72.110.94:3128 | ✓ 1958ms | 否 | ✓ 815ms | ✓ 1898ms | 否 | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1869ms | ✓ 1175ms | ✓ 766ms | http |
| 94.177.131.12:3128 | ✓ 856ms | 否 | ✓ 657ms | ✓ 1014ms | ✓ 778ms | http |
| 5.129.228.225:1080 | ✓ 601ms | 否 | 否 | ✓ 1178ms | ✓ 905ms | http |
| 115.76.5.32:10008 | ✓ 1931ms | 否 | 否 | ✓ 1761ms | ✓ 1575ms | http |
| 183.128.208.68:7890 | ✓ 1846ms | ✓ 1085ms | ✓ 1227ms | ✓ 1117ms | ✓ 895ms | http |
| 39.104.201.40:7890 | ✓ 1098ms | ✓ 1420ms | ✓ 1201ms | ✓ 1499ms | ✓ 1138ms | http |
| 193.181.35.26:8118 | ✓ 880ms | 否 | ✓ 1405ms | ✓ 1626ms | ✓ 1430ms | http |
| 45.136.198.40:3128 | ✓ 877ms | ✓ 1767ms | ✓ 578ms | 否 | 否 | http |
| 103.236.64.247:8888 | ✓ 1268ms | ✓ 1917ms | 否 | 否 | ✓ 1403ms | http |
| 36.147.78.166:80 | ✓ 1876ms | 否 | 否 | ✓ 1891ms | ✓ 1864ms | http |
| 85.198.84.77:10808 | ✓ 1223ms | 否 | ✓ 1839ms | 否 | ✓ 1542ms | http |
| 45.140.147.155:1082 | 否 | ✓ 1240ms | 否 | ✓ 1684ms | ✓ 1036ms | http |
| 45.129.141.143:3128 | ✓ 1368ms | 否 | ✓ 1677ms | ✓ 1794ms | ✓ 1754ms | http |
| 91.233.223.147:3128 | ✓ 920ms | 否 | ✓ 1904ms | 否 | ✓ 1927ms | http |
| 57.128.188.167:9173 | ✓ 1380ms | 否 | ✓ 1870ms | 否 | ✓ 1883ms | http |
| 222.28.182.229:7890 | 否 | ✓ 1494ms | ✓ 1515ms | ✓ 1534ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1820ms | 否 | ✓ 1021ms | ✓ 1488ms | ✓ 1173ms | http |
| 103.39.51.190:8080 | ✓ 1852ms | 否 | 否 | ✓ 1758ms | ✓ 1643ms | http |
| 125.128.12.144:3128 | ✓ 830ms | 否 | ✓ 1740ms | ✓ 1981ms | 否 | http |
| 61.72.221.234:3128 | 否 | ✓ 1729ms | ✓ 1293ms | 否 | ✓ 1687ms | http |
| 150.107.140.238:3128 | ✓ 1933ms | 否 | 否 | ✓ 1945ms | ✓ 1701ms | http |
| 171.234.62.116:10009 | 否 | 否 | ✓ 1861ms | ✓ 1711ms | ✓ 1441ms | http |
| 46.249.103.192:443 | ✓ 1141ms | 否 | ✓ 1368ms | ✓ 1929ms | 否 | http |
| 103.35.188.243:3128 | 否 | ✓ 1695ms | 否 | ✓ 1089ms | ✓ 796ms | http |
| 113.176.92.71:3128 | ✓ 1374ms | ✓ 1559ms | ✓ 1298ms | ✓ 1749ms | ✓ 1098ms | http |
| 121.230.9.205:1080 | 否 | ✓ 1525ms | 否 | ✓ 1801ms | ✓ 1344ms | http |
| 115.231.181.40:8128 | ✓ 1620ms | 否 | ✓ 1281ms | 否 | ✓ 1921ms | http |
| 36.147.78.166:443 | 否 | ✓ 1885ms | 否 | ✓ 1862ms | ✓ 1910ms | http |
| 2.56.178.131:443 | ✓ 1177ms | 否 | ✓ 1753ms | 否 | ✓ 1754ms | http |

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
