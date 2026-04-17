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

最后更新：2026-04-17 10:16:34 UTC（2026-04-17 18:16:34 UTC+8）

**代理总数：71**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 20.127.128.70:8080 | ✓ 826ms | 否 | ✓ 1258ms | ✓ 1988ms | ✓ 1255ms | http |
| 1.231.81.166:3128 | ✓ 1805ms | 否 | ✓ 1773ms | ✓ 1720ms | ✓ 1159ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1500ms | ✓ 1811ms | ✓ 1215ms | http |
| 218.108.131.186:17890 | ✓ 1012ms | ✓ 1208ms | ✓ 1069ms | ✓ 1468ms | ✓ 993ms | http |
| 188.246.224.49:7890 | ✓ 803ms | 否 | ✓ 579ms | 否 | ✓ 1597ms | http |
| 8.219.195.129:1080 | ✓ 1026ms | ✓ 1930ms | ✓ 991ms | ✓ 1354ms | ✓ 1055ms | http |
| 177.93.132.244:3128 | 否 | 否 | ✓ 699ms | ✓ 1977ms | ✓ 1517ms | http |
| 212.58.132.5:8888 | ✓ 1036ms | 否 | ✓ 1425ms | ✓ 1571ms | ✓ 1181ms | http |
| 185.138.116.150:8080 | 否 | 否 | ✓ 1998ms | ✓ 1643ms | ✓ 1128ms | http |
| 120.92.212.16:7890 | ✓ 1410ms | 否 | ✓ 1194ms | 否 | ✓ 1156ms | http |
| 117.236.124.166:3128 | ✓ 1596ms | 否 | ✓ 1779ms | 否 | ✓ 1949ms | http |
| 149.51.42.10:3128 | ✓ 656ms | ✓ 1519ms | 否 | ✓ 1643ms | 否 | http |
| 149.51.42.10:8080 | ✓ 1070ms | ✓ 1876ms | 否 | ✓ 1702ms | 否 | http |
| 36.141.21.200:7890 | ✓ 1031ms | ✓ 1341ms | ✓ 1083ms | ✓ 1344ms | ✓ 1069ms | http |
| 45.12.151.226:2829 | ✓ 784ms | ✓ 1751ms | ✓ 865ms | 否 | 否 | http |
| 120.92.108.86:7890 | ✓ 1432ms | 否 | ✓ 1580ms | 否 | ✓ 1641ms | http |
| 108.131.109.106:1516 | ✓ 1146ms | 否 | ✓ 1735ms | 否 | ✓ 1929ms | http |
| 185.114.73.2:1080 | ✓ 451ms | ✓ 1543ms | 否 | ✓ 1647ms | ✓ 1423ms | http |
| 38.242.223.121:3128 | ✓ 723ms | ✓ 1674ms | ✓ 1951ms | ✓ 1915ms | ✓ 1428ms | http |
| 202.141.161.53:10808 | ✓ 1081ms | ✓ 1259ms | ✓ 1296ms | 否 | ✓ 1236ms | http |
| 94.131.118.129:1081 | ✓ 710ms | ✓ 1570ms | ✓ 627ms | ✓ 1903ms | ✓ 1082ms | http |
| 27.71.24.102:3128 | ✓ 920ms | 否 | 否 | ✓ 1251ms | ✓ 981ms | http |
| 168.144.75.9:3128 | ✓ 1466ms | 否 | ✓ 1260ms | ✓ 1937ms | ✓ 1156ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1774ms | ✓ 1960ms | 否 | ✓ 1120ms | http |
| 45.140.147.155:1082 | ✓ 1178ms | 否 | ✓ 601ms | ✓ 1996ms | ✓ 1087ms | http |
| 45.140.147.155:1081 | ✓ 1158ms | ✓ 1116ms | ✓ 1485ms | 否 | ✓ 1084ms | http |
| 130.61.30.221:8080 | ✓ 905ms | 否 | ✓ 1311ms | 否 | ✓ 1538ms | http |
| 103.113.70.189:1081 | ✓ 581ms | ✓ 1306ms | ✓ 1894ms | 否 | 否 | http |
| 185.40.7.206:3128 | 否 | ✓ 1575ms | ✓ 1435ms | 否 | ✓ 1531ms | http |
| 51.145.178.158:3128 | ✓ 671ms | 否 | ✓ 1189ms | 否 | ✓ 1596ms | http |
| 149.104.4.88:10809 | ✓ 1751ms | ✓ 1842ms | ✓ 1300ms | 否 | ✓ 1173ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1655ms | ✓ 1366ms | ✓ 1957ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1680ms | 否 | ✓ 933ms | ✓ 1539ms | ✓ 1162ms | http |
| 82.114.228.67:1080 | ✓ 681ms | 否 | ✓ 1209ms | ✓ 1592ms | 否 | http |
| 103.113.70.189:1082 | ✓ 1662ms | ✓ 1027ms | ✓ 135ms | 否 | 否 | http |
| 84.47.150.125:1080 | ✓ 747ms | 否 | ✓ 928ms | 否 | ✓ 1724ms | http |
| 47.74.226.8:5001 | 否 | 否 | ✓ 1382ms | ✓ 1593ms | ✓ 1468ms | http |
| 101.32.243.189:80 | 否 | ✓ 1435ms | 否 | ✓ 1611ms | ✓ 1507ms | http |
| 15.161.131.175:3129 | ✓ 1824ms | 否 | ✓ 1766ms | 否 | ✓ 1666ms | http |
| 117.86.6.35:1080 | ✓ 1069ms | 否 | ✓ 1108ms | ✓ 1395ms | ✓ 1194ms | http |
| 139.159.97.82:10900 | ✓ 1973ms | 否 | 否 | ✓ 1661ms | ✓ 1416ms | http |
| 147.45.60.34:1082 | ✓ 1087ms | ✓ 1365ms | ✓ 132ms | 否 | ✓ 851ms | http |
| 124.243.150.41:3128 | 否 | 否 | ✓ 1281ms | ✓ 1190ms | ✓ 969ms | http |
| 62.113.119.14:8080 | ✓ 1453ms | 否 | ✓ 689ms | ✓ 1876ms | 否 | http |
| 85.215.50.161:80 | 否 | ✓ 1504ms | ✓ 1613ms | ✓ 1826ms | 否 | http |
| 34.71.229.255:3128 | 否 | 否 | ✓ 721ms | ✓ 1253ms | ✓ 885ms | http |
| 16.62.127.160:3128 | ✓ 1252ms | 否 | ✓ 1320ms | ✓ 1639ms | ✓ 1266ms | http |
| 103.85.113.66:9999 | ✓ 1281ms | ✓ 1814ms | ✓ 1398ms | ✓ 1961ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1192ms | 否 | ✓ 1432ms | ✓ 1712ms | ✓ 1411ms | http |
| 84.47.150.126:1080 | ✓ 1817ms | ✓ 1692ms | ✓ 1338ms | 否 | ✓ 1890ms | http |
| 116.171.106.78:3443 | 否 | ✓ 1658ms | ✓ 1634ms | 否 | ✓ 1975ms | http |
| 139.227.17.70:17890 | ✓ 1866ms | 否 | ✓ 1507ms | ✓ 1907ms | ✓ 1086ms | http |
| 116.58.161.203:26021 | ✓ 1892ms | 否 | ✓ 1207ms | ✓ 1582ms | 否 | http |
| 223.84.151.86:30005 | ✓ 1577ms | ✓ 1627ms | ✓ 1292ms | 否 | ✓ 1952ms | http |
| 35.225.22.61:80 | ✓ 359ms | ✓ 1743ms | ✓ 935ms | ✓ 1026ms | ✓ 920ms | http |
| 144.31.27.49:1080 | ✓ 751ms | 否 | ✓ 1375ms | 否 | ✓ 1788ms | http |
| 34.246.183.20:1111 | ✓ 1461ms | 否 | ✓ 802ms | ✓ 1583ms | 否 | http |
| 103.138.70.165:3129 | 否 | 否 | ✓ 1750ms | ✓ 1582ms | ✓ 1512ms | http |
| 47.93.216.160:1081 | ✓ 1591ms | ✓ 1900ms | ✓ 1508ms | 否 | ✓ 1112ms | http |
| 104.248.211.46:7890 | ✓ 616ms | ✓ 1843ms | 否 | ✓ 1526ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1153ms | ✓ 1561ms | ✓ 1853ms | 否 | 否 | http |
| 52.59.218.12:3128 | ✓ 590ms | 否 | ✓ 911ms | 否 | ✓ 1197ms | http |
| 2.27.18.184:1080 | 否 | 否 | ✓ 1863ms | ✓ 1762ms | ✓ 1666ms | http |
| 157.230.178.216:8088 | ✓ 619ms | 否 | ✓ 1135ms | 否 | ✓ 1726ms | http |
| 152.32.132.190:7890 | ✓ 1675ms | 否 | ✓ 818ms | ✓ 986ms | ✓ 814ms | http |
| 115.231.181.40:8128 | ✓ 1051ms | 否 | ✓ 1952ms | 否 | ✓ 1773ms | http |
| 138.124.99.216:8888 | ✓ 644ms | ✓ 1841ms | 否 | ✓ 1446ms | ✓ 1777ms | http |
| 61.52.131.172:8443 | ✓ 1045ms | ✓ 1364ms | ✓ 1150ms | ✓ 1435ms | ✓ 1093ms | http |
| 66.228.47.125:110 | ✓ 956ms | 否 | ✓ 1005ms | 否 | ✓ 1151ms | http |
| 133.18.123.225:26021 | ✓ 1785ms | ✓ 1551ms | ✓ 1644ms | ✓ 1508ms | ✓ 1674ms | http |
| 160.250.5.22:1 | ✓ 1522ms | 否 | ✓ 1551ms | ✓ 1516ms | ✓ 1518ms | http |

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
