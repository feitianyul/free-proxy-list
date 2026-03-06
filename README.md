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

最后更新：2026-03-06 15:35:47 UTC（2026-03-06 23:35:47 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 234ms | ✓ 1866ms | ✓ 753ms | ✓ 1184ms | ✓ 845ms | http |
| 136.49.39.94:8888 | ✓ 693ms | ✓ 1512ms | ✓ 671ms | 否 | ✓ 1038ms | http |
| 1.231.81.166:3128 | ✓ 907ms | ✓ 1388ms | ✓ 964ms | ✓ 1022ms | ✓ 818ms | http |
| 152.42.195.165:8888 | ✓ 844ms | 否 | ✓ 1709ms | ✓ 1640ms | ✓ 992ms | http |
| 47.101.149.27:9010 | ✓ 1495ms | 否 | ✓ 1512ms | 否 | ✓ 1243ms | http |
| 167.172.69.123:8080 | ✓ 883ms | 否 | ✓ 1319ms | ✓ 1190ms | ✓ 1254ms | http |
| 211.171.114.154:3128 | ✓ 1284ms | 否 | ✓ 1747ms | ✓ 1596ms | ✓ 1415ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1699ms | ✓ 1328ms | ✓ 1599ms | ✓ 1032ms | http |
| 35.225.22.61:80 | ✓ 400ms | 否 | ✓ 416ms | ✓ 1184ms | ✓ 787ms | http |
| 61.72.110.54:3128 | 否 | ✓ 1900ms | ✓ 1547ms | ✓ 1203ms | 否 | http |
| 125.128.12.144:3128 | ✓ 1282ms | 否 | ✓ 1065ms | 否 | ✓ 1602ms | http |
| 85.9.195.140:1080 | ✓ 914ms | 否 | ✓ 1861ms | 否 | ✓ 1214ms | http |
| 91.193.240.157:9877 | ✓ 968ms | 否 | ✓ 857ms | 否 | ✓ 1421ms | http |
| 121.128.121.54:3128 | ✓ 1349ms | 否 | ✓ 702ms | 否 | ✓ 925ms | http |
| 120.92.212.16:7890 | ✓ 1062ms | ✓ 1378ms | ✓ 1167ms | ✓ 1388ms | 否 | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1346ms | ✓ 1321ms | ✓ 1082ms | http |
| 120.92.212.16:8890 | ✓ 1084ms | ✓ 1386ms | ✓ 1133ms | 否 | 否 | http |
| 125.128.12.14:3128 | 否 | 否 | ✓ 1409ms | ✓ 1548ms | ✓ 1979ms | http |
| 81.70.169.194:80 | 否 | ✓ 1991ms | 否 | ✓ 1416ms | ✓ 1204ms | http |
| 103.84.95.54:7890 | ✓ 785ms | 否 | 否 | ✓ 977ms | ✓ 788ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1191ms | 否 | ✓ 1233ms | ✓ 909ms | http |
| 101.43.255.96:80 | ✓ 1178ms | 否 | ✓ 1046ms | ✓ 1407ms | 否 | http |
| 61.72.110.94:3128 | ✓ 1320ms | 否 | ✓ 1939ms | ✓ 1214ms | ✓ 1116ms | http |
| 14.56.107.244:3128 | ✓ 747ms | ✓ 1976ms | ✓ 866ms | ✓ 1199ms | 否 | http |
| 165.227.5.10:8888 | 否 | ✓ 1992ms | ✓ 1470ms | 否 | ✓ 696ms | http |
| 121.126.185.63:25152 | ✓ 1760ms | 否 | ✓ 1737ms | 否 | ✓ 1851ms | http |
| 61.72.221.194:3128 | ✓ 1377ms | ✓ 1110ms | 否 | ✓ 1787ms | ✓ 1364ms | http |
| 159.89.31.62:8080 | ✓ 459ms | 否 | 否 | ✓ 1894ms | ✓ 1585ms | http |
| 193.108.118.190:8888 | ✓ 541ms | 否 | ✓ 1579ms | ✓ 1664ms | ✓ 1703ms | http |
| 42.115.72.27:2064 | ✓ 1899ms | 否 | ✓ 1661ms | ✓ 1868ms | ✓ 1631ms | http |
| 178.236.245.59:3128 | ✓ 804ms | 否 | ✓ 1537ms | 否 | ✓ 1328ms | http |
| 61.72.221.94:3128 | ✓ 1779ms | 否 | ✓ 885ms | ✓ 1737ms | 否 | http |
| 178.236.245.17:3128 | ✓ 813ms | 否 | ✓ 1798ms | ✓ 1983ms | ✓ 1494ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1366ms | ✓ 1905ms | ✓ 1374ms | http |
| 23.94.182.50:12345 | ✓ 278ms | ✓ 1280ms | ✓ 850ms | ✓ 1294ms | ✓ 894ms | http |
| 106.14.203.63:3333 | ✓ 970ms | ✓ 1653ms | ✓ 1318ms | ✓ 1223ms | ✓ 970ms | http |
| 185.243.218.43:49153 | ✓ 1796ms | 否 | ✓ 1542ms | 否 | ✓ 1486ms | http |
| 106.14.205.114:483 | ✓ 1168ms | ✓ 1351ms | 否 | ✓ 1921ms | ✓ 1377ms | http |
| 154.37.208.132:30000 | ✓ 938ms | 否 | 否 | ✓ 1561ms | ✓ 1385ms | http |
| 202.129.206.239:3128 | ✓ 1551ms | 否 | 否 | ✓ 1854ms | ✓ 1899ms | http |
| 185.191.236.162:3128 | ✓ 1346ms | ✓ 1537ms | ✓ 1076ms | 否 | ✓ 1042ms | http |
| 172.212.68.37:3128 | ✓ 1465ms | ✓ 1630ms | 否 | ✓ 1144ms | ✓ 1145ms | http |
| 115.231.181.40:8128 | ✓ 1285ms | 否 | ✓ 1141ms | 否 | ✓ 1921ms | http |
| 167.172.69.123:80 | ✓ 1242ms | 否 | ✓ 1266ms | ✓ 1187ms | 否 | http |
| 42.115.72.27:2049 | ✓ 1884ms | 否 | ✓ 1694ms | ✓ 1908ms | 否 | http |
| 14.56.177.44:3128 | ✓ 1781ms | ✓ 1997ms | ✓ 1025ms | ✓ 1421ms | ✓ 1100ms | http |
| 42.115.72.27:2039 | ✓ 1865ms | 否 | ✓ 1616ms | ✓ 1842ms | ✓ 1633ms | http |
| 61.72.221.234:3128 | ✓ 1791ms | 否 | ✓ 855ms | ✓ 1625ms | ✓ 1170ms | http |
| 168.235.110.63:3128 | ✓ 1737ms | 否 | ✓ 1622ms | ✓ 1774ms | ✓ 1434ms | http |
| 103.139.138.194:3128 | ✓ 1888ms | 否 | ✓ 1116ms | ✓ 1611ms | ✓ 1162ms | http |
| 45.136.198.40:3128 | ✓ 687ms | ✓ 1545ms | ✓ 1584ms | ✓ 1892ms | ✓ 1611ms | http |
| 103.215.36.88:18272 | ✓ 1081ms | 否 | ✓ 1363ms | 否 | ✓ 1176ms | http |
| 103.104.99.29:80 | ✓ 1916ms | 否 | ✓ 1590ms | ✓ 1707ms | ✓ 1876ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1506ms | ✓ 1656ms | ✓ 1618ms | http |
| 103.39.51.190:8080 | ✓ 1898ms | 否 | ✓ 1378ms | ✓ 1542ms | 否 | http |
| 162.248.165.72:1080 | ✓ 1101ms | 否 | ✓ 1645ms | 否 | ✓ 1816ms | http |
| 159.223.42.219:3128 | ✓ 1316ms | 否 | 否 | ✓ 1443ms | ✓ 1333ms | http |
| 192.71.213.85:9091 | ✓ 1056ms | 否 | ✓ 1306ms | ✓ 1747ms | 否 | http |
| 192.71.213.85:9812 | ✓ 481ms | 否 | ✓ 474ms | ✓ 1535ms | 否 | http |
| 103.74.192.243:7890 | 否 | 否 | ✓ 1085ms | ✓ 1573ms | ✓ 1937ms | http |
| 221.127.195.224:8888 | ✓ 1278ms | 否 | ✓ 1279ms | 否 | ✓ 1273ms | http |
| 46.249.103.192:443 | ✓ 1183ms | 否 | ✓ 1058ms | ✓ 1760ms | 否 | http |
| 45.140.147.155:1081 | ✓ 453ms | 否 | ✓ 1361ms | ✓ 1708ms | ✓ 1284ms | http |
| 138.124.53.25:7443 | ✓ 594ms | ✓ 1810ms | ✓ 1529ms | 否 | 否 | http |
| 88.80.150.82:8080 | ✓ 770ms | ✓ 1713ms | ✓ 795ms | 否 | ✓ 1534ms | https |

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
