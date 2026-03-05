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

最后更新：2026-03-05 09:42:39 UTC（2026-03-05 17:42:39 UTC+8）

**代理总数：68**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | 否 | 否 | ✓ 893ms | ✓ 1291ms | ✓ 811ms | http |
| 181.78.79.155:999 | 否 | 否 | ✓ 1667ms | ✓ 1960ms | ✓ 1592ms | http |
| 94.177.131.12:3128 | ✓ 684ms | 否 | ✓ 1391ms | ✓ 1333ms | ✓ 835ms | http |
| 125.128.12.144:3128 | 否 | ✓ 1331ms | ✓ 1129ms | ✓ 1251ms | 否 | http |
| 103.84.95.54:7890 | ✓ 1052ms | 否 | ✓ 1050ms | ✓ 1321ms | 否 | http |
| 90.84.188.97:8000 | 否 | ✓ 1649ms | ✓ 1586ms | 否 | ✓ 1395ms | http |
| 14.56.107.244:3128 | ✓ 867ms | 否 | ✓ 1874ms | ✓ 1930ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1618ms | 否 | ✓ 774ms | ✓ 1161ms | ✓ 877ms | http |
| 192.166.82.55:1080 | ✓ 1097ms | ✓ 1417ms | 否 | ✓ 1341ms | ✓ 1593ms | http |
| 61.72.110.94:3128 | ✓ 1617ms | 否 | ✓ 1024ms | 否 | ✓ 992ms | http |
| 35.225.22.61:80 | ✓ 954ms | ✓ 1055ms | ✓ 365ms | ✓ 1026ms | ✓ 925ms | http |
| 120.92.212.16:7890 | ✓ 1369ms | ✓ 1720ms | ✓ 1222ms | 否 | ✓ 1181ms | http |
| 120.92.212.16:8890 | ✓ 1432ms | 否 | ✓ 1415ms | ✓ 1689ms | ✓ 1421ms | http |
| 138.124.53.25:7443 | ✓ 1236ms | 否 | ✓ 1913ms | 否 | ✓ 1504ms | http |
| 61.72.110.54:3128 | ✓ 1431ms | 否 | 否 | ✓ 1293ms | ✓ 1037ms | http |
| 125.128.12.14:3128 | ✓ 1423ms | ✓ 1491ms | 否 | ✓ 1573ms | 否 | http |
| 121.128.121.54:3128 | 否 | ✓ 1353ms | 否 | ✓ 1260ms | ✓ 991ms | http |
| 188.132.141.249:443 | ✓ 955ms | 否 | ✓ 1532ms | 否 | ✓ 1887ms | http |
| 165.227.5.10:8888 | ✓ 813ms | ✓ 1025ms | ✓ 950ms | 否 | ✓ 880ms | http |
| 160.238.65.4:3128 | ✓ 1055ms | 否 | ✓ 1436ms | 否 | ✓ 1636ms | http |
| 160.238.65.8:3128 | ✓ 1055ms | 否 | ✓ 1437ms | 否 | ✓ 1676ms | http |
| 160.238.65.3:3128 | ✓ 1057ms | 否 | ✓ 1432ms | 否 | ✓ 1680ms | http |
| 160.238.65.9:3128 | ✓ 1052ms | 否 | ✓ 1451ms | ✓ 1998ms | ✓ 1629ms | http |
| 160.238.65.2:3128 | ✓ 1051ms | 否 | ✓ 1440ms | 否 | ✓ 1682ms | http |
| 121.237.181.137:8888 | ✓ 1022ms | 否 | ✓ 1094ms | ✓ 1518ms | ✓ 1303ms | http |
| 91.193.240.157:9877 | ✓ 1092ms | 否 | ✓ 1454ms | 否 | ✓ 1601ms | http |
| 46.249.103.192:443 | ✓ 1595ms | 否 | ✓ 1262ms | ✓ 1866ms | 否 | http |
| 81.70.169.194:80 | ✓ 1175ms | ✓ 1513ms | ✓ 1741ms | ✓ 1575ms | ✓ 1450ms | http |
| 186.116.148.52:8080 | ✓ 1109ms | 否 | 否 | ✓ 1986ms | ✓ 1597ms | http |
| 160.238.65.7:3128 | ✓ 663ms | ✓ 1280ms | ✓ 454ms | ✓ 1227ms | ✓ 937ms | http |
| 160.238.65.6:3128 | ✓ 662ms | ✓ 1163ms | ✓ 480ms | ✓ 1218ms | ✓ 939ms | http |
| 160.238.65.5:3128 | ✓ 662ms | 否 | ✓ 384ms | 否 | ✓ 929ms | http |
| 185.191.236.162:3128 | ✓ 953ms | 否 | ✓ 1997ms | ✓ 1591ms | ✓ 1619ms | http |
| 116.80.82.230:3172 | ✓ 1661ms | 否 | ✓ 1941ms | 否 | ✓ 1786ms | http |
| 116.80.82.227:3172 | ✓ 1664ms | 否 | ✓ 1942ms | 否 | ✓ 1813ms | http |
| 168.235.110.63:3128 | 否 | 否 | ✓ 817ms | ✓ 1247ms | ✓ 737ms | http |
| 183.237.195.130:3128 | ✓ 1287ms | ✓ 1931ms | ✓ 1678ms | 否 | ✓ 1388ms | http |
| 101.43.255.96:80 | ✓ 1115ms | 否 | ✓ 1797ms | 否 | ✓ 1163ms | http |
| 35.234.17.221:8080 | ✓ 1521ms | 否 | 否 | ✓ 1250ms | ✓ 1386ms | http |
| 211.171.114.154:3128 | ✓ 1283ms | 否 | 否 | ✓ 1912ms | ✓ 1429ms | http |
| 45.140.147.155:1082 | ✓ 1852ms | 否 | ✓ 1832ms | 否 | ✓ 1456ms | http |
| 45.140.147.155:1081 | ✓ 1851ms | 否 | ✓ 1818ms | 否 | ✓ 1469ms | http |
| 8.219.97.248:80 | ✓ 1998ms | 否 | 否 | ✓ 1737ms | ✓ 1504ms | http |
| 5.75.196.26:40000 | ✓ 884ms | ✓ 1258ms | ✓ 498ms | 否 | 否 | http |
| 185.243.218.43:49153 | 否 | 否 | ✓ 1809ms | ✓ 1868ms | ✓ 1424ms | http |
| 45.136.198.40:3128 | ✓ 650ms | ✓ 1844ms | ✓ 1674ms | 否 | ✓ 1762ms | http |
| 61.72.221.194:3128 | ✓ 1405ms | ✓ 1244ms | 否 | ✓ 1910ms | 否 | http |
| 193.228.139.78:8888 | ✓ 1207ms | ✓ 1555ms | 否 | ✓ 1760ms | 否 | http |
| 207.254.71.62:8088 | ✓ 1652ms | 否 | ✓ 1519ms | 否 | ✓ 1458ms | http |
| 222.228.171.92:8080 | ✓ 1103ms | 否 | ✓ 1521ms | ✓ 1721ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1765ms | 否 | ✓ 1105ms | 否 | ✓ 1086ms | http |
| 89.185.85.138:1080 | ✓ 803ms | ✓ 1404ms | ✓ 1585ms | ✓ 1470ms | ✓ 1247ms | http |
| 114.214.162.73:26001 | ✓ 1301ms | ✓ 1959ms | ✓ 1623ms | ✓ 1591ms | 否 | http |
| 103.113.70.189:1081 | 否 | ✓ 1012ms | 否 | ✓ 1086ms | ✓ 847ms | http |
| 77.83.203.6:443 | ✓ 577ms | ✓ 1563ms | 否 | 否 | ✓ 1643ms | http |
| 77.83.203.5:443 | ✓ 1693ms | ✓ 1370ms | ✓ 1179ms | 否 | ✓ 1510ms | http |
| 74.48.78.224:2080 | ✓ 488ms | 否 | ✓ 608ms | ✓ 1135ms | 否 | http |
| 15.204.233.75:3128 | ✓ 483ms | 否 | ✓ 1611ms | ✓ 1063ms | ✓ 1282ms | http |
| 103.215.36.88:16541 | 否 | ✓ 1657ms | ✓ 1301ms | ✓ 1525ms | ✓ 1196ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1530ms | ✓ 1242ms | ✓ 1487ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1285ms | 否 | ✓ 720ms | ✓ 1653ms | ✓ 1197ms | http |
| 95.85.252.153:21064 | ✓ 503ms | ✓ 1386ms | ✓ 1279ms | 否 | 否 | http |
| 45.205.28.107:8080 | 否 | ✓ 1637ms | ✓ 977ms | ✓ 1161ms | ✓ 807ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1784ms | ✓ 1792ms | ✓ 1791ms | http |
| 103.39.51.190:8080 | ✓ 1652ms | 否 | ✓ 1993ms | ✓ 1894ms | 否 | http |
| 121.230.8.136:1080 | ✓ 1536ms | 否 | 否 | ✓ 1970ms | ✓ 1596ms | http |
| 172.212.68.37:3128 | 否 | ✓ 1398ms | ✓ 694ms | ✓ 1124ms | ✓ 1221ms | http |
| 88.80.150.82:8080 | ✓ 765ms | 否 | ✓ 697ms | ✓ 1625ms | 否 | https |

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
