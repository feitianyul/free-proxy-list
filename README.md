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

最后更新：2026-05-06 17:35:19 UTC（2026-05-07 01:35:19 UTC+8）

**代理总数：52**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | 否 | ✓ 1678ms | ✓ 1741ms | 否 | ✓ 1203ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1368ms | ✓ 1428ms | 否 | ✓ 1182ms | http |
| 120.92.212.16:7890 | ✓ 1064ms | ✓ 1373ms | ✓ 1872ms | ✓ 1400ms | ✓ 1102ms | http |
| 62.113.119.14:8080 | ✓ 1708ms | 否 | ✓ 1259ms | 否 | ✓ 1846ms | http |
| 84.47.150.125:1080 | ✓ 1679ms | 否 | ✓ 1945ms | 否 | ✓ 1566ms | http |
| 45.125.67.37:8443 | ✓ 1099ms | 否 | ✓ 1296ms | ✓ 1366ms | 否 | http |
| 14.143.222.113:10155 | ✓ 1625ms | 否 | ✓ 1217ms | ✓ 1568ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1063ms | ✓ 1357ms | ✓ 1482ms | ✓ 1421ms | ✓ 1200ms | http |
| 80.92.204.47:1081 | ✓ 985ms | 否 | ✓ 1943ms | ✓ 1918ms | ✓ 1218ms | http |
| 158.160.215.167:8124 | ✓ 1288ms | 否 | ✓ 1725ms | 否 | ✓ 1696ms | http |
| 20.164.75.153:8080 | ✓ 1364ms | 否 | ✓ 1181ms | 否 | ✓ 1716ms | http |
| 47.77.216.82:1080 | ✓ 318ms | ✓ 1157ms | ✓ 516ms | ✓ 1106ms | ✓ 772ms | http |
| 43.156.132.113:3128 | ✓ 1010ms | ✓ 1902ms | ✓ 951ms | ✓ 1281ms | ✓ 1005ms | http |
| 52.16.215.4:8320 | ✓ 1119ms | 否 | ✓ 1661ms | ✓ 1798ms | ✓ 1472ms | http |
| 116.118.48.147:3128 | 否 | 否 | ✓ 1466ms | ✓ 1420ms | ✓ 1340ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1533ms | ✓ 1468ms | ✓ 1415ms | 否 | http |
| 86.104.74.110:1081 | ✓ 1353ms | ✓ 1223ms | ✓ 563ms | ✓ 1223ms | ✓ 1835ms | http |
| 86.104.74.110:1082 | ✓ 1348ms | ✓ 1197ms | ✓ 592ms | ✓ 1199ms | ✓ 1062ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1545ms | ✓ 1063ms | ✓ 1316ms | ✓ 1074ms | http |
| 139.159.97.82:10900 | 否 | 否 | ✓ 1501ms | ✓ 1720ms | ✓ 1452ms | http |
| 47.112.25.109:7890 | 否 | ✓ 1331ms | 否 | ✓ 1697ms | ✓ 1354ms | http |
| 91.233.223.147:3128 | ✓ 1732ms | ✓ 1803ms | ✓ 1143ms | ✓ 1898ms | ✓ 1546ms | http |
| 193.160.209.58:1080 | ✓ 1723ms | 否 | 否 | ✓ 1798ms | ✓ 1606ms | http |
| 59.46.216.131:30001 | ✓ 1584ms | 否 | ✓ 1292ms | ✓ 1668ms | ✓ 1233ms | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 1120ms | ✓ 1789ms | ✓ 1434ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1411ms | ✓ 1471ms | ✓ 1499ms | http |
| 206.206.126.177:2412 | ✓ 1343ms | 否 | ✓ 1402ms | ✓ 1921ms | ✓ 1356ms | http |
| 77.110.119.136:3128 | ✓ 913ms | ✓ 1891ms | 否 | ✓ 1402ms | ✓ 1502ms | http |
| 1.231.81.166:3128 | ✓ 1822ms | ✓ 1927ms | 否 | 否 | ✓ 1127ms | http |
| 104.243.46.122:3128 | ✓ 798ms | 否 | ✓ 1333ms | ✓ 1249ms | ✓ 1778ms | http |
| 207.254.71.62:8088 | ✓ 1467ms | ✓ 1991ms | ✓ 1495ms | ✓ 1780ms | ✓ 1384ms | http |
| 113.176.92.71:3128 | ✓ 1723ms | 否 | ✓ 1498ms | ✓ 1435ms | ✓ 1302ms | http |
| 190.12.150.244:999 | ✓ 967ms | ✓ 1558ms | ✓ 1098ms | 否 | 否 | http |
| 120.92.108.86:7890 | ✓ 1360ms | 否 | ✓ 1430ms | 否 | ✓ 1992ms | http |
| 3.101.133.120:80 | ✓ 568ms | ✓ 1628ms | ✓ 1005ms | ✓ 1378ms | ✓ 1166ms | http |
| 91.217.81.131:1080 | ✓ 950ms | ✓ 1756ms | ✓ 1327ms | 否 | 否 | http |
| 45.63.88.46:1080 | ✓ 1320ms | 否 | 否 | ✓ 1297ms | ✓ 1704ms | http |
| 77.110.107.80:1080 | ✓ 1858ms | 否 | ✓ 713ms | ✓ 1542ms | ✓ 1965ms | http |
| 62.133.60.126:24558 | 否 | 否 | ✓ 1410ms | ✓ 1313ms | ✓ 1208ms | http |
| 15.161.131.175:20234 | 否 | 否 | ✓ 1595ms | ✓ 1533ms | ✓ 1241ms | http |
| 212.58.132.5:8888 | ✓ 1797ms | 否 | ✓ 1683ms | ✓ 1675ms | ✓ 1270ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1301ms | ✓ 1481ms | ✓ 1355ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1868ms | ✓ 1628ms | ✓ 1728ms | http |
| 223.16.170.103:80 | ✓ 1381ms | ✓ 1891ms | ✓ 1373ms | ✓ 1362ms | ✓ 1376ms | http |
| 103.18.77.14:1111 | 否 | 否 | ✓ 1851ms | ✓ 1601ms | ✓ 1657ms | http |
| 138.197.68.35:4857 | ✓ 123ms | ✓ 1239ms | 否 | 否 | ✓ 680ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1071ms | ✓ 295ms | ✓ 1663ms | 否 | http |
| 121.230.8.136:1080 | 否 | ✓ 1602ms | ✓ 1412ms | ✓ 1818ms | 否 | http |
| 101.32.243.189:80 | 否 | 否 | ✓ 1278ms | ✓ 1706ms | ✓ 1530ms | http |
| 61.52.131.172:8443 | ✓ 1597ms | 否 | ✓ 1143ms | 否 | ✓ 1166ms | http |
| 82.114.228.67:1080 | ✓ 721ms | ✓ 1441ms | ✓ 1264ms | ✓ 1501ms | 否 | http |
| 94.131.118.39:1081 | ✓ 1522ms | ✓ 1402ms | ✓ 944ms | ✓ 1618ms | ✓ 1159ms | http |

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
