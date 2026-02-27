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

最后更新：2026-02-27 17:41:52 UTC（2026-02-28 01:41:52 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 255ms | 否 | ✓ 1169ms | ✓ 1138ms | ✓ 1056ms | http |
| 104.238.30.45:59741 | ✓ 1702ms | 否 | ✓ 1839ms | 否 | ✓ 1967ms | http |
| 104.238.30.68:63744 | ✓ 1751ms | 否 | ✓ 1807ms | 否 | ✓ 1967ms | http |
| 104.238.30.63:63744 | ✓ 1751ms | 否 | ✓ 1775ms | 否 | ✓ 1968ms | http |
| 61.72.110.24:3128 | ✓ 1907ms | 否 | 否 | ✓ 1849ms | ✓ 989ms | http |
| 147.45.216.148:1080 | ✓ 477ms | 否 | 否 | ✓ 1706ms | ✓ 1645ms | http |
| 72.56.59.56:63127 | ✓ 1507ms | 否 | ✓ 1614ms | 否 | ✓ 1846ms | http |
| 104.238.30.39:59741 | ✓ 1695ms | 否 | ✓ 1775ms | 否 | ✓ 1967ms | http |
| 104.238.30.86:63900 | ✓ 1715ms | 否 | ✓ 1775ms | 否 | ✓ 1968ms | http |
| 104.238.30.50:59741 | ✓ 1678ms | 否 | ✓ 1807ms | 否 | ✓ 1967ms | http |
| 104.238.30.58:63744 | ✓ 1698ms | 否 | ✓ 1743ms | 否 | ✓ 1967ms | http |
| 195.123.209.48:3128 | ✓ 1253ms | 否 | ✓ 1347ms | ✓ 1900ms | ✓ 1419ms | http |
| 210.223.44.230:3128 | ✓ 1838ms | 否 | ✓ 1610ms | ✓ 1691ms | ✓ 1967ms | http |
| 104.238.30.37:59741 | ✓ 1729ms | 否 | ✓ 1807ms | 否 | ✓ 1999ms | http |
| 35.225.22.61:80 | ✓ 853ms | 否 | ✓ 908ms | 否 | ✓ 1147ms | http |
| 175.215.54.252:3040 | ✓ 1786ms | ✓ 1700ms | 否 | ✓ 1344ms | ✓ 960ms | http |
| 175.215.54.228:3128 | ✓ 1803ms | 否 | 否 | ✓ 1295ms | ✓ 917ms | http |
| 52.188.28.218:3128 | ✓ 1086ms | 否 | 否 | ✓ 907ms | ✓ 1714ms | http |
| 104.238.30.38:59741 | ✓ 1719ms | 否 | ✓ 1807ms | 否 | ✓ 1999ms | http |
| 81.177.48.54:2080 | ✓ 1248ms | 否 | ✓ 1515ms | ✓ 1934ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1045ms | ✓ 1350ms | ✓ 1126ms | ✓ 1381ms | ✓ 1064ms | http |
| 81.70.169.194:80 | ✓ 1151ms | ✓ 1350ms | ✓ 1474ms | ✓ 1411ms | ✓ 1178ms | http |
| 101.47.73.135:3128 | ✓ 1759ms | 否 | ✓ 1828ms | ✓ 1509ms | ✓ 1303ms | http |
| 121.237.181.137:8888 | ✓ 1403ms | ✓ 1521ms | ✓ 1986ms | ✓ 1905ms | ✓ 1007ms | http |
| 72.56.50.17:59787 | ✓ 1977ms | 否 | ✓ 1616ms | 否 | ✓ 1937ms | http |
| 35.234.17.221:8080 | ✓ 981ms | 否 | ✓ 1165ms | 否 | ✓ 1244ms | http |
| 120.46.152.136:3128 | ✓ 1494ms | ✓ 1960ms | 否 | ✓ 1822ms | 否 | http |
| 8.219.97.248:80 | ✓ 1746ms | 否 | ✓ 859ms | ✓ 1683ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1355ms | ✓ 1351ms | ✓ 1527ms | ✓ 1326ms | ✓ 1046ms | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 1128ms | ✓ 1082ms | ✓ 1886ms | http |
| 101.43.255.96:80 | ✓ 1320ms | ✓ 1408ms | 否 | ✓ 1321ms | ✓ 1124ms | http |
| 132.145.93.138:1080 | ✓ 1232ms | 否 | 否 | ✓ 1960ms | ✓ 1170ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1220ms | ✓ 996ms | 否 | ✓ 1811ms | http |
| 207.248.113.12:8080 | ✓ 720ms | ✓ 1811ms | 否 | ✓ 1430ms | ✓ 1242ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 819ms | ✓ 1660ms | ✓ 777ms | http |
| 72.56.59.62:63133 | ✓ 1561ms | 否 | ✓ 1619ms | 否 | ✓ 1903ms | http |
| 168.235.110.63:3128 | ✓ 290ms | 否 | 否 | ✓ 1093ms | ✓ 775ms | http |
| 165.227.5.10:8888 | ✓ 1618ms | 否 | ✓ 1131ms | ✓ 1280ms | ✓ 1130ms | http |
| 36.147.78.166:80 | 否 | ✓ 1817ms | ✓ 1806ms | 否 | ✓ 1777ms | http |
| 121.230.8.80:1080 | 否 | 否 | ✓ 1271ms | ✓ 1578ms | ✓ 1231ms | http |
| 36.147.78.166:443 | ✓ 1822ms | ✓ 1827ms | ✓ 1807ms | ✓ 1797ms | ✓ 1747ms | http |
| 59.46.216.131:30001 | ✓ 1028ms | 否 | ✓ 1390ms | 否 | ✓ 1147ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1658ms | ✓ 1626ms | ✓ 1606ms | http |
| 202.129.206.239:3128 | ✓ 1797ms | 否 | ✓ 1689ms | 否 | ✓ 1508ms | http |
| 104.238.30.91:63900 | ✓ 1755ms | 否 | ✓ 1808ms | 否 | ✓ 1999ms | http |
| 45.136.198.40:3128 | 否 | 否 | ✓ 1596ms | ✓ 1992ms | ✓ 1562ms | http |
| 103.82.23.118:5171 | 否 | 否 | ✓ 1203ms | ✓ 1766ms | ✓ 1403ms | http |
| 103.82.23.118:5234 | 否 | 否 | ✓ 1487ms | ✓ 1657ms | ✓ 1338ms | http |
| 34.205.52.219:80 | ✓ 330ms | 否 | ✓ 1471ms | ✓ 1821ms | ✓ 1891ms | http |
| 100.51.83.120:80 | ✓ 234ms | 否 | ✓ 515ms | ✓ 1771ms | 否 | http |
| 98.88.69.176:80 | ✓ 681ms | 否 | ✓ 687ms | ✓ 1886ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1680ms | ✓ 1544ms | ✓ 1586ms | ✓ 1589ms | ✓ 1150ms | http |
| 103.39.51.190:8080 | ✓ 1891ms | 否 | 否 | ✓ 1622ms | ✓ 1491ms | http |
| 62.113.119.14:8080 | ✓ 683ms | 否 | ✓ 897ms | 否 | ✓ 1552ms | http |
| 150.107.140.238:3128 | ✓ 947ms | 否 | ✓ 986ms | ✓ 1438ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1577ms | 否 | ✓ 1571ms | ✓ 1590ms | ✓ 1334ms | http |
| 103.143.197.41:8080 | ✓ 1899ms | 否 | ✓ 1801ms | ✓ 1881ms | ✓ 1629ms | http |
| 172.212.68.37:3128 | ✓ 326ms | ✓ 1622ms | ✓ 1060ms | ✓ 1200ms | ✓ 1080ms | http |
| 103.104.99.29:80 | ✓ 1960ms | 否 | ✓ 1805ms | ✓ 1683ms | ✓ 1789ms | http |
| 138.124.53.25:7443 | ✓ 1319ms | 否 | ✓ 1985ms | 否 | ✓ 1616ms | http |
| 45.186.6.104:3128 | ✓ 1127ms | ✓ 1633ms | ✓ 1983ms | 否 | 否 | http |
| 34.96.238.40:8080 | 否 | ✓ 1410ms | 否 | ✓ 1089ms | ✓ 959ms | http |

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
