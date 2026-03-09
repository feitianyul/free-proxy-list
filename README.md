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

最后更新：2026-03-09 23:29:09 UTC（2026-03-10 07:29:09 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.47:8443 | ✓ 915ms | ✓ 1141ms | ✓ 792ms | ✓ 1213ms | ✓ 761ms | http |
| 154.3.236.202:3128 | ✓ 321ms | 否 | ✓ 808ms | ✓ 1353ms | ✓ 1355ms | http |
| 162.240.154.26:3128 | ✓ 716ms | 否 | ✓ 1536ms | ✓ 1550ms | ✓ 1246ms | http |
| 1.231.81.166:3128 | ✓ 1792ms | ✓ 1119ms | ✓ 1100ms | ✓ 1118ms | ✓ 874ms | http |
| 91.107.141.42:8081 | ✓ 1175ms | ✓ 1543ms | ✓ 877ms | 否 | ✓ 1432ms | http |
| 61.72.110.114:3128 | ✓ 1794ms | 否 | ✓ 1463ms | 否 | ✓ 1174ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1566ms | ✓ 1768ms | ✓ 1101ms | http |
| 1.225.116.115:1080 | ✓ 1286ms | ✓ 1534ms | 否 | ✓ 1508ms | ✓ 1545ms | http |
| 35.225.22.61:80 | ✓ 936ms | 否 | 否 | ✓ 1139ms | ✓ 1076ms | http |
| 89.251.9.11:3128 | ✓ 385ms | 否 | ✓ 825ms | ✓ 1216ms | ✓ 1159ms | http |
| 185.191.236.162:3128 | ✓ 1085ms | ✓ 1811ms | ✓ 1218ms | 否 | ✓ 1498ms | http |
| 205.209.118.30:3138 | ✓ 798ms | 否 | ✓ 57ms | ✓ 1056ms | 否 | http |
| 43.225.185.4:8000 | ✓ 1912ms | 否 | ✓ 1371ms | ✓ 1607ms | ✓ 1545ms | http |
| 61.72.221.94:3128 | ✓ 1693ms | ✓ 1363ms | 否 | 否 | ✓ 999ms | http |
| 116.80.49.167:3172 | ✓ 1661ms | 否 | ✓ 1619ms | 否 | ✓ 1802ms | http |
| 116.80.96.102:3172 | 否 | 否 | ✓ 1686ms | ✓ 1992ms | ✓ 1781ms | http |
| 46.183.25.8:443 | ✓ 904ms | 否 | 否 | ✓ 1827ms | ✓ 1308ms | http |
| 194.213.18.200:443 | ✓ 1121ms | 否 | 否 | ✓ 1803ms | ✓ 1774ms | http |
| 202.155.12.161:443 | ✓ 1827ms | 否 | ✓ 1389ms | ✓ 1403ms | ✓ 1295ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1481ms | ✓ 1373ms | 否 | ✓ 1211ms | http |
| 120.92.212.16:7890 | ✓ 1521ms | ✓ 1500ms | ✓ 1173ms | 否 | ✓ 1201ms | http |
| 115.231.181.40:8128 | ✓ 1200ms | ✓ 1262ms | ✓ 1483ms | ✓ 1404ms | ✓ 1121ms | http |
| 81.70.169.194:80 | ✓ 1207ms | ✓ 1472ms | ✓ 1142ms | ✓ 1503ms | ✓ 1178ms | http |
| 101.43.255.96:80 | ✓ 1165ms | ✓ 1537ms | ✓ 1313ms | ✓ 1443ms | ✓ 1212ms | http |
| 190.9.109.198:999 | ✓ 694ms | ✓ 1326ms | ✓ 1235ms | ✓ 1262ms | ✓ 984ms | http |
| 190.9.109.207:999 | ✓ 689ms | ✓ 1331ms | ✓ 1234ms | ✓ 1249ms | 否 | http |
| 193.168.173.136:443 | ✓ 751ms | 否 | ✓ 1019ms | ✓ 1687ms | ✓ 1424ms | http |
| 210.223.44.230:3128 | ✓ 1854ms | 否 | ✓ 1920ms | ✓ 1380ms | ✓ 1150ms | http |
| 165.227.5.10:8888 | ✓ 577ms | 否 | 否 | ✓ 1004ms | ✓ 1407ms | http |
| 121.237.181.137:8888 | ✓ 1022ms | ✓ 1365ms | ✓ 1045ms | ✓ 1347ms | ✓ 1059ms | http |
| 192.71.213.85:9091 | ✓ 1256ms | 否 | ✓ 1500ms | ✓ 1905ms | 否 | http |
| 152.42.213.210:8080 | ✓ 1912ms | 否 | ✓ 1910ms | 否 | ✓ 1350ms | http |
| 61.72.110.54:3128 | ✓ 985ms | ✓ 1463ms | ✓ 1758ms | 否 | ✓ 1851ms | http |
| 116.80.49.156:3172 | ✓ 1718ms | 否 | ✓ 1718ms | 否 | ✓ 1859ms | http |
| 45.186.6.104:3128 | ✓ 1322ms | ✓ 1778ms | ✓ 1606ms | 否 | 否 | http |
| 116.80.49.165:3172 | ✓ 1830ms | 否 | ✓ 1789ms | 否 | ✓ 1815ms | http |
| 94.176.3.43:7443 | ✓ 991ms | 否 | ✓ 1549ms | 否 | ✓ 1671ms | http |
| 136.49.34.18:8888 | ✓ 1438ms | 否 | ✓ 1700ms | ✓ 1890ms | 否 | http |
| 125.128.12.144:3128 | ✓ 1472ms | 否 | ✓ 1921ms | ✓ 1922ms | 否 | http |
| 103.35.188.243:3128 | ✓ 338ms | ✓ 1730ms | ✓ 790ms | ✓ 1156ms | ✓ 947ms | http |
| 210.77.29.245:7890 | ✓ 985ms | ✓ 1306ms | 否 | ✓ 1309ms | ✓ 1028ms | http |
| 103.183.10.172:3125 | 否 | 否 | ✓ 1883ms | ✓ 1972ms | ✓ 1645ms | http |
| 88.80.150.82:8080 | ✓ 1512ms | ✓ 1732ms | ✓ 1819ms | 否 | ✓ 1664ms | https |
| 83.219.250.8:62920 | ✓ 870ms | 否 | ✓ 1312ms | 否 | ✓ 1408ms | http |
| 103.183.10.169:3125 | ✓ 1975ms | 否 | 否 | ✓ 1700ms | ✓ 1681ms | http |
| 45.136.130.223:8443 | ✓ 875ms | ✓ 1505ms | ✓ 1254ms | ✓ 954ms | ✓ 790ms | http |
| 34.101.184.164:3128 | ✓ 1845ms | 否 | ✓ 1776ms | ✓ 1592ms | ✓ 1467ms | http |
| 168.235.110.63:3128 | ✓ 722ms | ✓ 1072ms | ✓ 1399ms | ✓ 1064ms | ✓ 763ms | http |
| 95.3.9.78:3128 | ✓ 966ms | 否 | 否 | ✓ 1585ms | ✓ 1188ms | http |
| 14.225.212.37:7890 | ✓ 1677ms | ✓ 1692ms | ✓ 1740ms | ✓ 1325ms | ✓ 1028ms | http |
| 154.53.40.110:3128 | ✓ 416ms | 否 | ✓ 1586ms | ✓ 1442ms | ✓ 764ms | http |
| 43.167.227.161:1080 | ✓ 705ms | 否 | ✓ 704ms | ✓ 956ms | ✓ 782ms | http |
| 185.254.53.154:8080 | ✓ 1014ms | ✓ 1776ms | ✓ 1620ms | 否 | ✓ 1510ms | http |
| 116.58.162.45:3128 | ✓ 1401ms | ✓ 1675ms | ✓ 929ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 1131ms | 否 | 否 | ✓ 1866ms | ✓ 1665ms | http |
| 172.212.68.37:3128 | ✓ 574ms | 否 | ✓ 983ms | ✓ 1416ms | ✓ 982ms | http |
| 138.124.53.25:7443 | ✓ 472ms | ✓ 1924ms | ✓ 1134ms | ✓ 1782ms | 否 | http |
| 45.129.141.143:3128 | ✓ 896ms | 否 | ✓ 1717ms | ✓ 1786ms | ✓ 1589ms | http |
| 39.104.201.40:7890 | ✓ 1087ms | 否 | ✓ 1081ms | ✓ 1496ms | ✓ 1104ms | http |
| 47.101.159.19:8899 | ✓ 1039ms | ✓ 1260ms | ✓ 1782ms | ✓ 1386ms | ✓ 1061ms | http |
| 201.144.20.238:3128 | ✓ 900ms | 否 | ✓ 1252ms | ✓ 1242ms | ✓ 1071ms | http |
| 45.140.147.82:1081 | ✓ 575ms | 否 | ✓ 413ms | 否 | ✓ 1348ms | http |
| 125.128.12.14:3128 | ✓ 790ms | ✓ 1091ms | ✓ 1395ms | ✓ 1643ms | ✓ 1370ms | http |
| 152.42.213.210:80 | ✓ 960ms | 否 | ✓ 929ms | ✓ 1611ms | ✓ 1103ms | http |
| 121.230.8.55:1080 | ✓ 1177ms | ✓ 1530ms | ✓ 1197ms | ✓ 1673ms | ✓ 1401ms | http |
| 113.177.131.2:3128 | ✓ 1806ms | 否 | ✓ 1243ms | 否 | ✓ 1479ms | http |
| 172.104.63.237:3128 | ✓ 1272ms | 否 | ✓ 907ms | ✓ 1948ms | ✓ 1469ms | http |
| 123.25.25.180:1452 | ✓ 1598ms | 否 | 否 | ✓ 1794ms | ✓ 1690ms | http |
| 152.70.98.46:8888 | ✓ 932ms | ✓ 1533ms | ✓ 1588ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1952ms | 否 | ✓ 1465ms | ✓ 1643ms | ✓ 1604ms | http |
| 43.165.195.107:3128 | ✓ 1887ms | 否 | ✓ 1557ms | ✓ 1435ms | 否 | http |
| 111.228.59.13:8888 | ✓ 1741ms | ✓ 1621ms | ✓ 1655ms | 否 | 否 | http |
| 103.183.10.203:3125 | ✓ 1550ms | 否 | ✓ 1570ms | ✓ 1847ms | ✓ 1621ms | http |

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
