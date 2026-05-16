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

最后更新：2026-05-16 17:52:31 UTC（2026-05-17 01:52:31 UTC+8）

**代理总数：61**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 170.106.136.181:31002 | ✓ 326ms | ✓ 828ms | ✓ 677ms | ✓ 1296ms | ✓ 580ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1813ms | ✓ 976ms | ✓ 1324ms | ✓ 1206ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1868ms | ✓ 1108ms | 否 | ✓ 1092ms | http |
| 185.200.188.234:10001 | ✓ 1133ms | 否 | ✓ 1885ms | 否 | ✓ 1790ms | http |
| 51.161.50.166:3128 | ✓ 510ms | ✓ 1786ms | 否 | ✓ 1719ms | ✓ 1251ms | http |
| 218.108.131.186:17890 | ✓ 892ms | ✓ 1063ms | ✓ 873ms | ✓ 1125ms | ✓ 932ms | http |
| 207.254.71.62:8088 | ✓ 1087ms | ✓ 1798ms | ✓ 816ms | ✓ 1550ms | ✓ 1396ms | http |
| 91.242.229.129:8092 | 否 | 否 | ✓ 1679ms | ✓ 1881ms | ✓ 1839ms | http |
| 8.154.21.175:3128 | ✓ 890ms | ✓ 1101ms | ✓ 895ms | ✓ 1180ms | ✓ 963ms | http |
| 114.214.170.41:27890 | ✓ 1143ms | ✓ 1420ms | ✓ 1359ms | ✓ 1470ms | ✓ 1194ms | http |
| 148.230.4.241:999 | ✓ 848ms | 否 | ✓ 760ms | ✓ 1566ms | ✓ 1258ms | http |
| 128.199.254.13:9090 | ✓ 767ms | 否 | ✓ 1226ms | ✓ 1791ms | 否 | http |
| 152.42.170.187:9090 | ✓ 1627ms | 否 | ✓ 1580ms | ✓ 1592ms | ✓ 1525ms | http |
| 146.190.80.158:9090 | ✓ 797ms | 否 | ✓ 1184ms | ✓ 1989ms | 否 | http |
| 57.129.144.178:40000 | ✓ 1086ms | 否 | ✓ 1210ms | ✓ 1761ms | ✓ 1599ms | http |
| 103.235.67.190:80 | 否 | 否 | ✓ 1223ms | ✓ 1278ms | ✓ 996ms | http |
| 158.255.212.55:3256 | ✓ 1627ms | 否 | ✓ 1657ms | ✓ 1892ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1173ms | ✓ 1956ms | ✓ 961ms | ✓ 1536ms | ✓ 1170ms | http |
| 194.59.247.34:10808 | ✓ 1166ms | ✓ 1657ms | ✓ 1457ms | 否 | ✓ 1729ms | http |
| 84.47.150.125:1080 | ✓ 1178ms | 否 | ✓ 1163ms | 否 | ✓ 1790ms | http |
| 3.15.187.17:1080 | ✓ 1228ms | 否 | ✓ 996ms | ✓ 1713ms | ✓ 1137ms | http |
| 121.130.199.80:24003 | 否 | ✓ 1011ms | 否 | ✓ 1496ms | ✓ 1009ms | http |
| 178.63.155.151:8888 | ✓ 1873ms | 否 | ✓ 1516ms | 否 | ✓ 1918ms | http |
| 120.92.212.16:7890 | ✓ 1333ms | ✓ 1185ms | ✓ 1525ms | ✓ 1545ms | ✓ 1034ms | http |
| 120.92.212.16:8890 | ✓ 1868ms | ✓ 1864ms | ✓ 950ms | 否 | ✓ 1037ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1497ms | ✓ 1449ms | ✓ 1939ms | http |
| 166.88.55.83:7890 | ✓ 782ms | ✓ 1214ms | ✓ 1817ms | ✓ 937ms | ✓ 728ms | http |
| 121.230.8.137:1080 | ✓ 1308ms | 否 | 否 | ✓ 1478ms | ✓ 1062ms | http |
| 150.107.140.238:3128 | ✓ 1744ms | 否 | ✓ 1150ms | ✓ 1173ms | 否 | http |
| 128.199.114.189:9090 | ✓ 1556ms | 否 | ✓ 1917ms | ✓ 1581ms | ✓ 976ms | http |
| 3.101.133.120:80 | ✓ 262ms | 否 | ✓ 1632ms | ✓ 1170ms | ✓ 1050ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1127ms | ✓ 1373ms | ✓ 1906ms | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1491ms | ✓ 1615ms | ✓ 1320ms | http |
| 210.223.44.230:3128 | ✓ 902ms | 否 | 否 | ✓ 1003ms | ✓ 1631ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1876ms | ✓ 1459ms | ✓ 1617ms | http |
| 115.231.181.40:8128 | ✓ 1159ms | 否 | ✓ 1975ms | 否 | ✓ 1064ms | http |
| 129.80.238.83:444 | ✓ 325ms | 否 | 否 | ✓ 1151ms | ✓ 885ms | http |
| 129.80.217.21:444 | ✓ 1630ms | 否 | ✓ 1344ms | 否 | ✓ 1344ms | http |
| 116.171.106.111:3443 | ✓ 1633ms | ✓ 1658ms | ✓ 1518ms | ✓ 1663ms | ✓ 1861ms | http |
| 43.156.90.221:10808 | 否 | 否 | ✓ 1802ms | ✓ 1039ms | ✓ 831ms | http |
| 158.255.212.55:9005 | ✓ 1029ms | 否 | ✓ 1922ms | ✓ 1894ms | 否 | http |
| 158.255.212.55:9480 | ✓ 1034ms | 否 | ✓ 1916ms | ✓ 1895ms | 否 | http |
| 158.255.212.55:7497 | ✓ 1109ms | 否 | ✓ 1847ms | ✓ 1885ms | 否 | http |
| 158.255.212.55:7839 | ✓ 1021ms | 否 | ✓ 1922ms | ✓ 1897ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1823ms | ✓ 1403ms | ✓ 1233ms | ✓ 1270ms | ✓ 984ms | http |
| 217.182.195.221:30003 | ✓ 1566ms | 否 | ✓ 1423ms | 否 | ✓ 1693ms | http |
| 113.176.92.71:3128 | ✓ 1519ms | 否 | ✓ 930ms | ✓ 1280ms | 否 | http |
| 106.10.55.212:1121 | ✓ 1672ms | ✓ 1898ms | ✓ 897ms | ✓ 1411ms | ✓ 1209ms | http |
| 8.219.97.248:80 | ✓ 1000ms | 否 | 否 | ✓ 1638ms | ✓ 1634ms | http |
| 43.167.192.85:8080 | ✓ 1323ms | 否 | ✓ 1544ms | ✓ 919ms | ✓ 835ms | http |
| 185.191.236.162:3128 | ✓ 859ms | 否 | ✓ 1728ms | ✓ 1838ms | ✓ 1229ms | http |
| 103.69.84.106:3131 | ✓ 1777ms | 否 | ✓ 1519ms | ✓ 1201ms | ✓ 966ms | http |
| 103.21.220.141:3128 | ✓ 899ms | 否 | ✓ 1486ms | ✓ 1211ms | ✓ 1958ms | http |
| 61.52.131.172:8443 | ✓ 883ms | ✓ 1218ms | ✓ 1064ms | ✓ 1247ms | ✓ 1004ms | http |
| 103.172.70.173:8080 | ✓ 1757ms | 否 | 否 | ✓ 1458ms | ✓ 1455ms | http |
| 212.58.132.5:8888 | ✓ 1915ms | 否 | ✓ 1977ms | ✓ 1696ms | 否 | http |
| 86.104.72.219:1081 | ✓ 1015ms | ✓ 1116ms | ✓ 912ms | ✓ 1479ms | ✓ 1040ms | http |
| 168.110.52.228:3128 | ✓ 1685ms | 否 | 否 | ✓ 845ms | ✓ 1679ms | http |
| 152.32.132.190:7890 | ✓ 1093ms | ✓ 1822ms | 否 | ✓ 1917ms | 否 | http |
| 121.230.8.136:1080 | ✓ 1954ms | 否 | ✓ 1655ms | ✓ 1392ms | ✓ 1822ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1740ms | ✓ 1882ms | ✓ 1890ms | http |

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
