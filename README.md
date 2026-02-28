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

最后更新：2026-02-28 08:35:11 UTC（2026-02-28 16:35:11 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 439ms | ✓ 1993ms | ✓ 1076ms | ✓ 1136ms | ✓ 834ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1732ms | ✓ 1331ms | ✓ 1994ms | http |
| 132.145.93.138:1080 | 否 | 否 | ✓ 1927ms | ✓ 1264ms | ✓ 1087ms | http |
| 52.188.28.218:3128 | ✓ 1239ms | 否 | ✓ 1993ms | ✓ 1293ms | 否 | http |
| 3.213.157.4:3128 | ✓ 196ms | ✓ 1495ms | ✓ 811ms | ✓ 1378ms | ✓ 1398ms | http |
| 120.202.127.234:10808 | ✓ 909ms | 否 | ✓ 1063ms | ✓ 1404ms | ✓ 1033ms | http |
| 103.84.95.54:7890 | ✓ 1113ms | 否 | ✓ 760ms | 否 | ✓ 989ms | http |
| 147.45.216.148:1080 | ✓ 492ms | 否 | ✓ 1394ms | ✓ 1824ms | ✓ 1286ms | http |
| 45.140.147.155:1081 | ✓ 1432ms | 否 | ✓ 1338ms | ✓ 1478ms | ✓ 1268ms | http |
| 104.238.30.58:63744 | ✓ 1721ms | 否 | ✓ 1807ms | 否 | ✓ 1967ms | http |
| 104.238.30.45:59741 | ✓ 1782ms | 否 | ✓ 1775ms | 否 | ✓ 1972ms | http |
| 120.92.212.16:7890 | ✓ 1436ms | ✓ 1328ms | ✓ 1120ms | ✓ 1382ms | ✓ 1361ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1220ms | ✓ 1801ms | ✓ 1292ms | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 1100ms | ✓ 1083ms | ✓ 885ms | http |
| 222.28.182.229:7890 | ✓ 1126ms | 否 | 否 | ✓ 1453ms | ✓ 1963ms | http |
| 120.92.212.16:8890 | ✓ 1247ms | ✓ 1332ms | ✓ 1903ms | ✓ 1331ms | 否 | http |
| 36.147.78.166:80 | ✓ 1859ms | 否 | ✓ 1801ms | ✓ 1924ms | 否 | http |
| 81.70.169.194:80 | ✓ 1532ms | ✓ 1421ms | ✓ 1200ms | ✓ 1345ms | ✓ 1886ms | http |
| 104.238.30.63:63744 | ✓ 1877ms | 否 | ✓ 1963ms | 否 | ✓ 1999ms | http |
| 47.101.149.27:9010 | ✓ 1444ms | ✓ 1396ms | 否 | 否 | ✓ 1240ms | http |
| 121.40.231.103:7890 | 否 | ✓ 1150ms | ✓ 970ms | ✓ 1255ms | ✓ 969ms | http |
| 121.148.239.82:3127 | 否 | 否 | ✓ 1436ms | ✓ 1420ms | ✓ 892ms | http |
| 34.101.184.164:3128 | ✓ 1718ms | 否 | ✓ 1343ms | ✓ 1505ms | ✓ 1509ms | http |
| 101.43.255.96:80 | ✓ 1201ms | ✓ 1490ms | ✓ 1083ms | ✓ 1441ms | ✓ 1382ms | http |
| 103.39.51.207:8080 | 否 | 否 | ✓ 1373ms | ✓ 1717ms | ✓ 1451ms | http |
| 85.208.108.43:2094 | ✓ 392ms | 否 | ✓ 228ms | ✓ 1181ms | ✓ 759ms | http |
| 168.235.110.63:3128 | ✓ 555ms | ✓ 1047ms | ✓ 1026ms | ✓ 1098ms | ✓ 773ms | http |
| 121.237.181.137:8888 | 否 | ✓ 1314ms | ✓ 1057ms | 否 | ✓ 1013ms | http |
| 35.234.17.221:8080 | ✓ 926ms | 否 | 否 | ✓ 1285ms | ✓ 955ms | http |
| 147.161.234.13:10613 | ✓ 475ms | 否 | ✓ 758ms | ✓ 1373ms | 否 | http |
| 147.161.234.13:11727 | ✓ 592ms | 否 | ✓ 730ms | ✓ 1508ms | 否 | http |
| 147.161.234.13:10919 | ✓ 582ms | 否 | ✓ 741ms | ✓ 1573ms | 否 | http |
| 165.225.73.113:10919 | ✓ 609ms | 否 | ✓ 623ms | ✓ 1337ms | ✓ 1104ms | http |
| 165.225.73.112:10919 | ✓ 664ms | 否 | ✓ 567ms | ✓ 1331ms | ✓ 1090ms | http |
| 165.225.80.216:12605 | ✓ 480ms | 否 | ✓ 1454ms | ✓ 1857ms | 否 | http |
| 165.225.73.113:11551 | ✓ 574ms | 否 | ✓ 749ms | ✓ 1513ms | ✓ 1068ms | http |
| 165.225.206.16:10919 | ✓ 595ms | ✓ 1762ms | ✓ 979ms | ✓ 1702ms | ✓ 1453ms | http |
| 165.225.206.12:10919 | ✓ 617ms | 否 | ✓ 976ms | ✓ 1957ms | ✓ 991ms | http |
| 165.225.21.228:12522 | ✓ 878ms | 否 | ✓ 930ms | ✓ 1750ms | ✓ 1242ms | http |
| 165.225.20.12:12183 | ✓ 876ms | 否 | ✓ 1394ms | ✓ 1369ms | ✓ 1287ms | http |
| 165.225.20.12:11011 | ✓ 877ms | 否 | ✓ 958ms | ✓ 1859ms | ✓ 1235ms | http |
| 165.225.204.14:11711 | ✓ 1086ms | ✓ 1578ms | ✓ 859ms | ✓ 1891ms | ✓ 1589ms | http |
| 165.225.21.227:11797 | ✓ 875ms | 否 | ✓ 1400ms | ✓ 1355ms | ✓ 1326ms | http |
| 165.225.21.227:12691 | ✓ 877ms | 否 | ✓ 1977ms | ✓ 1237ms | ✓ 949ms | http |
| 165.225.204.16:11416 | ✓ 878ms | ✓ 1524ms | ✓ 1119ms | ✓ 1994ms | ✓ 1591ms | http |
| 165.225.21.228:11153 | ✓ 876ms | 否 | ✓ 937ms | ✓ 1749ms | ✓ 1593ms | http |
| 165.225.80.214:11587 | ✓ 523ms | 否 | ✓ 1409ms | 否 | ✓ 1342ms | http |
| 165.225.206.16:12670 | ✓ 637ms | 否 | ✓ 720ms | ✓ 1861ms | 否 | http |
| 147.161.190.35:11129 | ✓ 868ms | 否 | ✓ 1573ms | ✓ 1690ms | 否 | http |
| 165.225.120.17:11178 | 否 | 否 | ✓ 1322ms | ✓ 1915ms | ✓ 1492ms | http |
| 165.225.120.17:10906 | 否 | 否 | ✓ 1630ms | ✓ 1786ms | ✓ 1362ms | http |
| 36.147.78.166:443 | ✓ 1796ms | ✓ 1515ms | 否 | ✓ 1985ms | ✓ 1738ms | http |
| 138.124.53.25:7443 | ✓ 1136ms | 否 | 否 | ✓ 1848ms | ✓ 1435ms | http |
| 52.201.29.25:80 | ✓ 172ms | 否 | ✓ 771ms | ✓ 1699ms | ✓ 929ms | http |
| 45.136.198.40:3128 | ✓ 1812ms | ✓ 1852ms | ✓ 1730ms | ✓ 1799ms | 否 | http |
| 98.95.199.207:80 | ✓ 108ms | ✓ 1242ms | ✓ 797ms | ✓ 1030ms | 否 | http |
| 100.52.6.187:80 | ✓ 786ms | 否 | ✓ 826ms | 否 | ✓ 1028ms | http |
| 3.225.78.45:80 | 否 | 否 | ✓ 797ms | ✓ 1835ms | ✓ 1267ms | http |
| 100.49.145.52:80 | ✓ 783ms | 否 | ✓ 1494ms | 否 | ✓ 1803ms | http |
| 150.107.140.238:3128 | ✓ 1945ms | 否 | ✓ 1086ms | 否 | ✓ 1112ms | http |
| 115.231.181.40:8128 | ✓ 1612ms | ✓ 1237ms | ✓ 1026ms | ✓ 1685ms | 否 | http |
| 101.255.107.34:8080 | ✓ 1407ms | 否 | 否 | ✓ 1516ms | ✓ 1506ms | http |
| 172.212.68.37:3128 | ✓ 337ms | 否 | ✓ 764ms | ✓ 960ms | ✓ 804ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1366ms | ✓ 1836ms | ✓ 1683ms | http |
| 62.113.119.14:8080 | ✓ 1130ms | 否 | ✓ 1408ms | ✓ 1624ms | ✓ 1106ms | http |
| 72.56.50.17:59787 | ✓ 1521ms | 否 | ✓ 1643ms | 否 | ✓ 1941ms | http |
| 43.161.214.161:1081 | 否 | 否 | ✓ 1685ms | ✓ 1905ms | ✓ 1542ms | http |
| 147.161.135.21:12362 | ✓ 1024ms | ✓ 1746ms | ✓ 1147ms | ✓ 1472ms | ✓ 1408ms | http |
| 147.161.135.21:11138 | ✓ 1024ms | ✓ 1746ms | ✓ 893ms | ✓ 1726ms | 否 | http |
| 165.225.80.40:12851 | ✓ 1027ms | 否 | ✓ 1238ms | ✓ 1624ms | ✓ 1887ms | http |
| 165.225.80.14:10919 | ✓ 1027ms | ✓ 1420ms | 否 | 否 | ✓ 1863ms | http |
| 167.103.6.24:12174 | ✓ 1813ms | 否 | ✓ 1907ms | 否 | ✓ 1872ms | http |
| 45.140.147.82:1081 | ✓ 865ms | ✓ 1978ms | 否 | 否 | ✓ 1320ms | http |
| 61.72.110.94:3128 | ✓ 748ms | 否 | ✓ 1233ms | 否 | ✓ 941ms | http |
| 45.140.147.82:1082 | ✓ 523ms | 否 | ✓ 1128ms | ✓ 1819ms | ✓ 1207ms | http |
| 103.236.64.247:8888 | ✓ 1357ms | 否 | 否 | ✓ 1979ms | ✓ 1557ms | http |

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
