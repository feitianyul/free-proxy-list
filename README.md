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

最后更新：2026-04-29 18:33:41 UTC（2026-04-30 02:33:41 UTC+8）

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
| 218.108.131.186:17890 | ✓ 991ms | ✓ 1276ms | ✓ 1086ms | ✓ 1328ms | ✓ 1070ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1603ms | ✓ 1629ms | ✓ 1380ms | ✓ 1029ms | http |
| 34.71.229.255:3128 | ✓ 1667ms | ✓ 1627ms | ✓ 1094ms | ✓ 1835ms | ✓ 1600ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1386ms | ✓ 1550ms | ✓ 1192ms | http |
| 115.231.181.40:8128 | ✓ 1106ms | 否 | ✓ 1067ms | 否 | ✓ 1940ms | http |
| 45.167.124.71:999 | ✓ 1731ms | 否 | ✓ 1290ms | ✓ 1573ms | ✓ 1429ms | http |
| 47.85.51.197:1080 | ✓ 106ms | ✓ 953ms | ✓ 1515ms | 否 | ✓ 664ms | http |
| 43.99.54.236:5555 | ✓ 1771ms | ✓ 1452ms | ✓ 861ms | ✓ 1096ms | ✓ 821ms | http |
| 46.101.95.183:8888 | ✓ 1092ms | 否 | ✓ 964ms | 否 | ✓ 1296ms | http |
| 154.64.232.35:8080 | ✓ 1367ms | ✓ 1654ms | ✓ 1365ms | 否 | 否 | http |
| 86.104.74.110:1081 | ✓ 1704ms | ✓ 1070ms | ✓ 766ms | 否 | 否 | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1216ms | ✓ 1498ms | ✓ 1296ms | http |
| 45.140.147.155:1082 | ✓ 1411ms | ✓ 1782ms | 否 | 否 | ✓ 1520ms | http |
| 62.113.119.14:8080 | ✓ 1328ms | ✓ 1393ms | ✓ 874ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1321ms | ✓ 1414ms | ✓ 1152ms | 否 | 否 | http |
| 3.238.34.111:3128 | ✓ 1222ms | 否 | ✓ 1835ms | ✓ 1948ms | ✓ 1755ms | http |
| 3.19.213.118:40000 | ✓ 1205ms | 否 | ✓ 1835ms | 否 | ✓ 1814ms | http |
| 15.237.108.20:47140 | ✓ 1213ms | 否 | ✓ 1716ms | 否 | ✓ 1716ms | http |
| 172.236.145.31:7890 | ✓ 1121ms | 否 | 否 | ✓ 1842ms | ✓ 1288ms | http |
| 77.110.116.224:3128 | 否 | 否 | ✓ 1051ms | ✓ 1745ms | ✓ 1365ms | http |
| 103.157.200.126:3128 | ✓ 1239ms | 否 | ✓ 1182ms | ✓ 1640ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1151ms | ✓ 1343ms | 否 | ✓ 1590ms | ✓ 1079ms | http |
| 152.70.91.193:40000 | ✓ 1503ms | 否 | 否 | ✓ 1647ms | ✓ 1233ms | http |
| 45.10.71.107:8888 | 否 | 否 | ✓ 1031ms | ✓ 1195ms | ✓ 1254ms | http |
| 212.58.132.5:8888 | ✓ 1132ms | 否 | ✓ 1418ms | ✓ 1728ms | ✓ 1177ms | http |
| 8.211.166.184:8081 | ✓ 1207ms | ✓ 1207ms | ✓ 913ms | ✓ 1125ms | ✓ 881ms | http |
| 38.180.192.119:3128 | ✓ 1895ms | 否 | ✓ 1122ms | ✓ 889ms | ✓ 730ms | http |
| 120.92.212.16:8890 | ✓ 1407ms | 否 | ✓ 1661ms | ✓ 1663ms | 否 | http |
| 8.154.21.175:3128 | ✓ 970ms | ✓ 1229ms | ✓ 1058ms | ✓ 1335ms | ✓ 1104ms | http |
| 176.100.39.18:8080 | ✓ 1264ms | 否 | ✓ 1753ms | ✓ 1909ms | 否 | http |
| 51.44.97.6:44066 | ✓ 731ms | 否 | ✓ 1172ms | 否 | ✓ 1280ms | http |
| 15.223.237.12:33606 | ✓ 911ms | 否 | ✓ 1831ms | 否 | ✓ 1277ms | http |
| 122.52.107.138:8082 | 否 | 否 | ✓ 1952ms | ✓ 1737ms | ✓ 1844ms | http |
| 78.12.252.87:23972 | ✓ 1230ms | 否 | ✓ 1271ms | ✓ 1932ms | 否 | http |
| 121.130.177.28:8888 | 否 | ✓ 1603ms | 否 | ✓ 1596ms | ✓ 1817ms | http |
| 194.87.26.83:3128 | ✓ 900ms | ✓ 1470ms | 否 | ✓ 1390ms | ✓ 1442ms | http |
| 103.189.249.133:1111 | ✓ 1984ms | 否 | 否 | ✓ 1736ms | ✓ 1766ms | http |
| 8.219.97.248:80 | ✓ 1632ms | 否 | ✓ 1484ms | ✓ 1705ms | 否 | http |
| 183.232.248.73:7890 | ✓ 1061ms | ✓ 1339ms | ✓ 1140ms | ✓ 1259ms | ✓ 1021ms | http |
| 52.47.115.41:10603 | ✓ 639ms | 否 | 否 | ✓ 1497ms | ✓ 1516ms | http |
| 34.85.118.216:3128 | 否 | 否 | ✓ 1332ms | ✓ 1235ms | ✓ 996ms | http |
| 45.76.207.177:40000 | 否 | 否 | ✓ 1366ms | ✓ 1550ms | ✓ 1188ms | http |
| 13.41.196.179:2897 | ✓ 698ms | 否 | ✓ 1804ms | 否 | ✓ 1627ms | http |
| 108.131.109.106:58080 | ✓ 1246ms | 否 | ✓ 1291ms | 否 | ✓ 1829ms | http |
| 78.12.252.87:47978 | ✓ 1020ms | 否 | ✓ 1757ms | 否 | ✓ 1670ms | http |
| 13.48.13.125:443 | ✓ 1021ms | 否 | ✓ 1274ms | ✓ 1989ms | 否 | http |
| 103.209.36.58:8080 | ✓ 1877ms | 否 | ✓ 1329ms | ✓ 1463ms | 否 | http |
| 20.27.15.111:8561 | 否 | 否 | ✓ 1057ms | ✓ 1076ms | ✓ 1779ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 1079ms | ✓ 1090ms | ✓ 1799ms | http |
| 20.27.13.35:8561 | 否 | 否 | ✓ 1052ms | ✓ 1104ms | ✓ 1786ms | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 1076ms | ✓ 1078ms | ✓ 1815ms | http |
| 45.140.147.82:1082 | ✓ 1103ms | ✓ 1257ms | ✓ 837ms | ✓ 1775ms | ✓ 1492ms | http |
| 45.140.147.82:1081 | ✓ 1112ms | ✓ 1283ms | ✓ 792ms | ✓ 1788ms | ✓ 1543ms | http |
| 59.46.216.131:30001 | ✓ 1197ms | 否 | 否 | ✓ 1650ms | ✓ 1274ms | http |
| 47.112.25.109:7890 | ✓ 1125ms | ✓ 1705ms | 否 | ✓ 1373ms | 否 | http |
| 77.110.107.80:1080 | ✓ 1179ms | 否 | ✓ 1602ms | ✓ 1641ms | ✓ 1119ms | http |
| 77.110.107.80:8080 | ✓ 1183ms | 否 | ✓ 1610ms | ✓ 1597ms | 否 | http |
| 34.96.238.40:8080 | 否 | ✓ 1338ms | ✓ 1267ms | ✓ 1399ms | 否 | http |
| 103.167.171.147:8080 | ✓ 1541ms | 否 | ✓ 1900ms | ✓ 1858ms | 否 | http |
| 101.32.243.189:80 | ✓ 1910ms | 否 | ✓ 1880ms | 否 | ✓ 1566ms | http |
| 61.52.131.172:8443 | ✓ 1131ms | ✓ 1346ms | ✓ 1171ms | ✓ 1368ms | ✓ 1144ms | http |
| 52.59.218.12:3534 | ✓ 1148ms | 否 | ✓ 1479ms | ✓ 1951ms | ✓ 1678ms | http |
| 103.39.51.207:8080 | ✓ 1580ms | 否 | 否 | ✓ 1956ms | ✓ 1782ms | http |

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
