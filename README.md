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

最后更新：2026-02-27 23:28:04 UTC（2026-02-28 07:28:04 UTC+8）

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
| 205.209.118.30:3138 | ✓ 376ms | 否 | ✓ 1171ms | ✓ 1379ms | ✓ 1040ms | http |
| 14.56.107.244:3128 | ✓ 1404ms | 否 | 否 | ✓ 932ms | ✓ 759ms | http |
| 3.213.157.4:3128 | ✓ 344ms | 否 | ✓ 1051ms | ✓ 1872ms | ✓ 1378ms | http |
| 138.124.53.25:7443 | ✓ 727ms | 否 | ✓ 1835ms | 否 | ✓ 1616ms | http |
| 45.140.147.82:1081 | ✓ 1629ms | ✓ 1633ms | ✓ 1466ms | 否 | 否 | http |
| 35.234.17.221:8080 | ✓ 798ms | ✓ 1221ms | ✓ 1076ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1066ms | ✓ 1303ms | 否 | ✓ 1283ms | 否 | http |
| 1.225.116.115:1080 | ✓ 1746ms | ✓ 1205ms | ✓ 721ms | ✓ 1016ms | ✓ 714ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1094ms | ✓ 888ms | ✓ 1061ms | ✓ 903ms | http |
| 101.47.73.135:3128 | ✓ 1062ms | 否 | ✓ 1571ms | ✓ 1075ms | ✓ 969ms | http |
| 61.72.110.24:3128 | ✓ 1744ms | ✓ 1805ms | ✓ 1014ms | ✓ 1378ms | 否 | http |
| 195.123.209.48:3128 | ✓ 1113ms | 否 | ✓ 1978ms | 否 | ✓ 1963ms | http |
| 121.237.181.137:8888 | 否 | 否 | ✓ 1840ms | ✓ 1144ms | ✓ 1836ms | http |
| 52.188.28.218:3128 | ✓ 397ms | ✓ 1766ms | 否 | 否 | ✓ 865ms | http |
| 181.78.194.249:999 | ✓ 1117ms | 否 | ✓ 1052ms | ✓ 1843ms | ✓ 1609ms | http |
| 35.225.22.61:80 | ✓ 588ms | 否 | 否 | ✓ 1207ms | ✓ 1093ms | http |
| 103.84.95.54:7890 | ✓ 635ms | 否 | 否 | ✓ 764ms | ✓ 839ms | http |
| 120.92.212.16:7890 | ✓ 1041ms | ✓ 1372ms | ✓ 1185ms | ✓ 1182ms | ✓ 1946ms | http |
| 61.72.110.94:3128 | ✓ 1666ms | ✓ 1321ms | ✓ 805ms | 否 | ✓ 1620ms | http |
| 101.43.255.96:80 | ✓ 967ms | ✓ 1176ms | ✓ 893ms | ✓ 1302ms | ✓ 944ms | http |
| 36.147.78.166:443 | 否 | ✓ 1570ms | 否 | ✓ 1531ms | ✓ 1546ms | http |
| 81.70.169.194:80 | ✓ 1009ms | ✓ 1302ms | ✓ 854ms | ✓ 1186ms | ✓ 1764ms | http |
| 36.147.78.166:80 | 否 | ✓ 1638ms | ✓ 1572ms | ✓ 1789ms | ✓ 1526ms | http |
| 45.125.67.37:8443 | ✓ 816ms | 否 | ✓ 868ms | ✓ 1127ms | 否 | http |
| 120.92.212.16:8890 | ✓ 925ms | ✓ 1162ms | ✓ 1185ms | ✓ 1183ms | ✓ 1155ms | http |
| 111.79.111.126:3128 | ✓ 1071ms | ✓ 1349ms | ✓ 1385ms | ✓ 1441ms | ✓ 1085ms | http |
| 103.236.64.247:8888 | ✓ 1349ms | 否 | ✓ 1258ms | 否 | ✓ 886ms | http |
| 220.197.44.36:3128 | 否 | ✓ 1754ms | ✓ 1870ms | 否 | ✓ 1648ms | http |
| 147.45.216.148:1080 | ✓ 679ms | 否 | ✓ 1248ms | 否 | ✓ 1304ms | http |
| 91.238.104.171:2023 | ✓ 883ms | 否 | ✓ 1251ms | 否 | ✓ 1341ms | http |
| 45.88.0.98:3128 | 否 | ✓ 1608ms | ✓ 1363ms | 否 | ✓ 1717ms | http |
| 168.235.110.63:3128 | ✓ 513ms | 否 | ✓ 971ms | ✓ 1286ms | ✓ 1068ms | http |
| 34.205.52.219:80 | ✓ 736ms | ✓ 1349ms | ✓ 1101ms | ✓ 1149ms | ✓ 896ms | http |
| 152.32.255.24:27197 | ✓ 1822ms | 否 | ✓ 1486ms | 否 | ✓ 1037ms | http |
| 61.72.110.54:3128 | ✓ 1415ms | ✓ 1441ms | ✓ 1863ms | 否 | 否 | http |
| 165.227.5.10:8888 | 否 | ✓ 814ms | 否 | ✓ 1038ms | ✓ 601ms | http |
| 120.46.152.136:3128 | ✓ 1032ms | ✓ 1334ms | ✓ 1455ms | ✓ 1624ms | ✓ 1074ms | http |
| 62.113.119.14:8080 | ✓ 1087ms | 否 | ✓ 1337ms | ✓ 1985ms | ✓ 1773ms | http |
| 45.88.0.115:3128 | ✓ 1191ms | 否 | ✓ 948ms | 否 | ✓ 1278ms | http |
| 116.80.80.173:3172 | 否 | 否 | ✓ 1641ms | ✓ 1778ms | ✓ 1626ms | http |
| 34.101.184.164:3128 | ✓ 1003ms | 否 | ✓ 1033ms | ✓ 1743ms | ✓ 1353ms | http |
| 175.106.14.126:3128 | ✓ 964ms | 否 | ✓ 1768ms | ✓ 1332ms | 否 | http |
| 113.176.92.71:3128 | ✓ 1983ms | ✓ 1395ms | ✓ 1419ms | ✓ 1229ms | ✓ 1309ms | http |
| 37.220.139.219:8080 | ✓ 1282ms | ✓ 1954ms | ✓ 1887ms | 否 | 否 | http |
| 85.208.108.43:2094 | ✓ 842ms | 否 | ✓ 1729ms | ✓ 1196ms | ✓ 997ms | http |
| 194.59.204.87:9080 | ✓ 1171ms | ✓ 1760ms | 否 | ✓ 1850ms | 否 | http |
| 94.177.131.12:3128 | ✓ 429ms | ✓ 1591ms | ✓ 586ms | ✓ 751ms | ✓ 840ms | http |
| 91.233.223.147:3128 | ✓ 1159ms | 否 | ✓ 945ms | 否 | ✓ 1657ms | http |
| 121.230.8.246:1080 | ✓ 1321ms | ✓ 1417ms | ✓ 1564ms | 否 | ✓ 1687ms | http |
| 162.240.154.26:3128 | ✓ 1071ms | ✓ 1395ms | ✓ 623ms | ✓ 1280ms | ✓ 1302ms | http |
| 103.82.23.118:5178 | ✓ 1505ms | ✓ 1739ms | ✓ 1385ms | 否 | ✓ 1483ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1833ms | ✓ 777ms | 否 | ✓ 837ms | http |
| 45.140.147.155:1081 | ✓ 672ms | 否 | ✓ 1496ms | ✓ 1792ms | ✓ 1379ms | http |
| 91.107.148.58:53967 | ✓ 630ms | 否 | 否 | ✓ 1743ms | ✓ 1010ms | http |
| 150.107.140.238:3128 | ✓ 1872ms | 否 | ✓ 1244ms | ✓ 1698ms | ✓ 1405ms | http |
| 8.219.97.248:80 | ✓ 877ms | 否 | ✓ 1769ms | ✓ 1283ms | 否 | http |
| 81.177.48.54:2080 | ✓ 1319ms | 否 | ✓ 1646ms | ✓ 1923ms | ✓ 1804ms | http |
| 172.212.68.37:3128 | ✓ 675ms | ✓ 1757ms | ✓ 1522ms | ✓ 1397ms | ✓ 1091ms | http |
| 45.136.198.40:3128 | 否 | ✓ 1996ms | ✓ 1797ms | 否 | ✓ 1943ms | http |
| 14.56.118.154:3128 | ✓ 1792ms | ✓ 1659ms | ✓ 1425ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 628ms | 否 | ✓ 692ms | ✓ 1476ms | ✓ 1176ms | http |
| 157.0.142.246:10057 | ✓ 1130ms | ✓ 1543ms | ✓ 969ms | 否 | ✓ 930ms | http |
| 49.146.153.161:8082 | ✓ 1396ms | 否 | 否 | ✓ 1382ms | ✓ 1378ms | http |
| 103.143.197.90:8080 | 否 | 否 | ✓ 1420ms | ✓ 1548ms | ✓ 1455ms | http |
| 103.39.51.190:8080 | ✓ 1597ms | 否 | ✓ 1682ms | ✓ 1301ms | ✓ 1485ms | http |
| 14.56.177.44:3128 | ✓ 1574ms | 否 | ✓ 1414ms | ✓ 1411ms | 否 | http |
| 121.128.121.224:3128 | 否 | ✓ 1704ms | ✓ 1811ms | 否 | ✓ 1900ms | http |
| 121.230.9.252:1080 | ✓ 1287ms | ✓ 1515ms | ✓ 1140ms | ✓ 1489ms | ✓ 1062ms | http |
| 103.104.99.29:80 | ✓ 1759ms | 否 | ✓ 1599ms | ✓ 1517ms | ✓ 1394ms | http |
| 103.104.99.89:80 | ✓ 1753ms | 否 | ✓ 1708ms | ✓ 1500ms | ✓ 1390ms | http |
| 91.238.104.172:2024 | ✓ 1741ms | ✓ 1877ms | ✓ 1561ms | 否 | ✓ 1921ms | http |
| 139.159.236.196:6868 | ✓ 1973ms | 否 | ✓ 1954ms | 否 | ✓ 1994ms | http |
| 121.230.9.26:1080 | ✓ 1846ms | 否 | ✓ 1111ms | ✓ 1918ms | 否 | http |
| 14.56.118.164:3128 | ✓ 1234ms | 否 | ✓ 1646ms | ✓ 1677ms | 否 | http |
| 120.28.220.41:8082 | ✓ 1271ms | 否 | ✓ 1693ms | ✓ 1428ms | ✓ 1433ms | http |

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
