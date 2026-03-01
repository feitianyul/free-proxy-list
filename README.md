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

最后更新：2026-03-01 14:43:30 UTC（2026-03-01 22:43:30 UTC+8）

**代理总数：66**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 403ms | 否 | ✓ 951ms | ✓ 1160ms | ✓ 1114ms | http |
| 62.113.119.14:8080 | ✓ 626ms | 否 | ✓ 741ms | ✓ 1753ms | ✓ 1294ms | http |
| 168.235.110.63:3128 | 否 | ✓ 1007ms | ✓ 999ms | ✓ 1790ms | 否 | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1130ms | ✓ 951ms | ✓ 774ms | http |
| 20.27.11.248:8561 | ✓ 987ms | ✓ 1004ms | ✓ 832ms | ✓ 981ms | ✓ 744ms | http |
| 148.135.85.87:1080 | ✓ 1292ms | ✓ 1955ms | 否 | ✓ 1482ms | ✓ 1346ms | http |
| 20.27.14.220:8561 | ✓ 1741ms | 否 | 否 | ✓ 983ms | ✓ 759ms | http |
| 20.27.15.111:8561 | ✓ 1754ms | 否 | 否 | ✓ 974ms | ✓ 777ms | http |
| 141.11.210.35:1080 | ✓ 466ms | 否 | ✓ 1033ms | 否 | ✓ 857ms | http |
| 39.104.201.40:7890 | ✓ 1023ms | ✓ 1339ms | ✓ 1122ms | ✓ 1348ms | ✓ 1070ms | http |
| 120.92.212.16:8890 | ✓ 1392ms | ✓ 1467ms | ✓ 1404ms | 否 | 否 | http |
| 167.160.184.231:6005 | ✓ 732ms | 否 | ✓ 782ms | ✓ 1019ms | ✓ 799ms | http |
| 5.75.201.136:1080 | ✓ 1054ms | 否 | 否 | ✓ 1635ms | ✓ 1417ms | http |
| 120.92.212.16:7890 | ✓ 1566ms | ✓ 1365ms | ✓ 1100ms | ✓ 1342ms | ✓ 1687ms | http |
| 101.43.255.96:80 | ✓ 1493ms | 否 | ✓ 1135ms | 否 | ✓ 1067ms | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 1313ms | ✓ 1432ms | ✓ 1089ms | http |
| 103.215.36.88:19328 | ✓ 1849ms | ✓ 1505ms | ✓ 1893ms | 否 | ✓ 1146ms | http |
| 59.46.216.131:30001 | ✓ 1563ms | 否 | ✓ 1319ms | ✓ 1438ms | 否 | http |
| 20.78.118.91:8561 | ✓ 1002ms | ✓ 1861ms | ✓ 877ms | ✓ 1038ms | ✓ 860ms | http |
| 20.210.39.153:8561 | ✓ 1007ms | 否 | ✓ 788ms | ✓ 1074ms | ✓ 837ms | http |
| 20.78.26.206:8561 | ✓ 1006ms | 否 | ✓ 790ms | ✓ 1084ms | ✓ 848ms | http |
| 165.227.5.10:8888 | ✓ 1656ms | 否 | ✓ 1484ms | 否 | ✓ 792ms | http |
| 74.208.234.198:443 | ✓ 735ms | ✓ 1683ms | ✓ 445ms | ✓ 1093ms | ✓ 672ms | http |
| 138.124.53.25:7443 | ✓ 1472ms | ✓ 1722ms | 否 | 否 | ✓ 1602ms | http |
| 121.128.121.54:3128 | ✓ 1202ms | 否 | ✓ 1328ms | ✓ 1171ms | ✓ 1182ms | http |
| 115.231.181.40:8128 | ✓ 1891ms | ✓ 1363ms | 否 | ✓ 1239ms | ✓ 1124ms | http |
| 20.27.15.49:8561 | 否 | 否 | ✓ 1691ms | ✓ 1391ms | ✓ 718ms | http |
| 20.210.76.104:8561 | 否 | 否 | ✓ 1692ms | ✓ 1416ms | ✓ 736ms | http |
| 35.234.17.221:8080 | ✓ 921ms | 否 | ✓ 1217ms | 否 | ✓ 1128ms | http |
| 210.223.44.230:3128 | ✓ 1400ms | ✓ 1398ms | ✓ 1006ms | ✓ 1055ms | ✓ 890ms | http |
| 185.243.218.43:49153 | ✓ 1242ms | ✓ 1603ms | ✓ 1355ms | ✓ 1990ms | ✓ 1519ms | http |
| 108.181.0.167:8080 | ✓ 1154ms | ✓ 1570ms | ✓ 244ms | ✓ 921ms | ✓ 802ms | http |
| 81.70.169.194:80 | 否 | ✓ 1316ms | ✓ 1679ms | 否 | ✓ 1119ms | http |
| 45.136.198.40:3128 | ✓ 1169ms | 否 | ✓ 1469ms | ✓ 1922ms | ✓ 1838ms | http |
| 85.208.108.43:10808 | ✓ 126ms | 否 | ✓ 1323ms | ✓ 1309ms | ✓ 888ms | http |
| 85.208.108.43:2094 | ✓ 127ms | 否 | ✓ 1343ms | ✓ 1322ms | ✓ 885ms | http |
| 14.56.107.244:3128 | 否 | ✓ 1005ms | ✓ 944ms | 否 | ✓ 1977ms | http |
| 104.238.30.58:63744 | ✓ 1762ms | 否 | ✓ 1839ms | 否 | ✓ 1995ms | http |
| 104.238.30.45:59741 | ✓ 1775ms | 否 | ✓ 1843ms | 否 | ✓ 1996ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1819ms | ✓ 1727ms | ✓ 1123ms | http |
| 14.56.177.44:3128 | ✓ 1104ms | ✓ 1684ms | ✓ 718ms | ✓ 1116ms | ✓ 1006ms | http |
| 121.128.121.184:3128 | ✓ 824ms | ✓ 1508ms | ✓ 1077ms | ✓ 1110ms | ✓ 1097ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1045ms | 否 | ✓ 1293ms | ✓ 898ms | http |
| 45.125.67.37:8443 | ✓ 1999ms | 否 | ✓ 1496ms | ✓ 1318ms | ✓ 1862ms | http |
| 45.140.147.82:1081 | ✓ 622ms | 否 | 否 | ✓ 1394ms | ✓ 1259ms | http |
| 104.248.25.131:3128 | ✓ 492ms | 否 | 否 | ✓ 1651ms | ✓ 1424ms | http |
| 110.172.29.131:3128 | ✓ 1800ms | 否 | ✓ 1776ms | ✓ 1352ms | ✓ 1092ms | http |
| 185.198.27.38:3128 | ✓ 519ms | ✓ 1500ms | ✓ 819ms | 否 | ✓ 1364ms | http |
| 104.238.30.50:59741 | ✓ 1717ms | 否 | ✓ 1776ms | 否 | ✓ 1999ms | http |
| 180.127.149.228:1080 | ✓ 1355ms | ✓ 1994ms | ✓ 1055ms | ✓ 1423ms | ✓ 1121ms | http |
| 77.83.203.5:443 | ✓ 1377ms | 否 | ✓ 595ms | ✓ 1437ms | ✓ 1050ms | http |
| 34.101.184.164:3128 | ✓ 1727ms | 否 | ✓ 992ms | ✓ 1372ms | ✓ 1093ms | http |
| 61.72.221.94:3128 | ✓ 1426ms | 否 | ✓ 1465ms | ✓ 1204ms | ✓ 939ms | http |
| 61.72.110.94:3128 | ✓ 1609ms | ✓ 1609ms | ✓ 1825ms | 否 | 否 | http |
| 61.72.110.54:3128 | ✓ 1993ms | ✓ 1064ms | 否 | ✓ 1151ms | 否 | http |
| 37.187.109.70:10111 | ✓ 1872ms | 否 | ✓ 1988ms | 否 | ✓ 1940ms | http |
| 172.212.68.37:3128 | ✓ 282ms | ✓ 1579ms | ✓ 1756ms | ✓ 1255ms | ✓ 1107ms | http |
| 95.85.252.153:21064 | ✓ 1371ms | ✓ 1864ms | ✓ 1482ms | 否 | 否 | http |
| 150.107.140.238:3128 | ✓ 1059ms | 否 | 否 | ✓ 1636ms | ✓ 1030ms | http |
| 103.82.23.118:5249 | ✓ 1309ms | 否 | ✓ 1475ms | 否 | ✓ 1622ms | http |
| 36.147.78.166:80 | ✓ 1780ms | ✓ 1782ms | 否 | 否 | ✓ 1855ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1949ms | ✓ 1658ms | ✓ 1626ms | http |
| 45.140.147.155:1082 | ✓ 439ms | 否 | ✓ 1821ms | 否 | ✓ 976ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1657ms | ✓ 1658ms | ✓ 1320ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1795ms | ✓ 1430ms | ✓ 1454ms | http |
| 36.147.78.166:443 | ✓ 1837ms | ✓ 1815ms | ✓ 1789ms | ✓ 1951ms | ✓ 1603ms | http |

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
