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

最后更新：2026-04-12 18:33:47 UTC（2026-04-13 02:33:47 UTC+8）

**代理总数：54**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 54 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 54 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 790ms | ✓ 1009ms | ✓ 808ms | ✓ 1037ms | ✓ 843ms | http |
| 36.141.21.200:7890 | ✓ 981ms | ✓ 1071ms | ✓ 940ms | ✓ 1176ms | ✓ 943ms | http |
| 147.161.210.140:8800 | ✓ 1498ms | 否 | ✓ 1123ms | ✓ 857ms | ✓ 672ms | http |
| 1.231.81.166:3128 | ✓ 1528ms | ✓ 1059ms | ✓ 1598ms | ✓ 1010ms | ✓ 879ms | http |
| 223.84.151.86:30005 | ✓ 1278ms | ✓ 1170ms | ✓ 1074ms | ✓ 1383ms | ✓ 1140ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1284ms | ✓ 909ms | ✓ 1052ms | ✓ 838ms | http |
| 167.103.115.102:8800 | ✓ 1088ms | 否 | ✓ 1001ms | ✓ 1268ms | ✓ 1193ms | http |
| 79.132.136.58:3128 | ✓ 820ms | 否 | 否 | ✓ 1493ms | ✓ 1152ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 1886ms | ✓ 1440ms | ✓ 1624ms | http |
| 35.225.22.61:80 | ✓ 790ms | 否 | ✓ 1466ms | ✓ 1596ms | ✓ 1106ms | http |
| 167.103.34.108:8800 | ✓ 1300ms | 否 | ✓ 1096ms | ✓ 1322ms | 否 | http |
| 36.103.198.235:7890 | ✓ 977ms | ✓ 1611ms | ✓ 1215ms | ✓ 1335ms | ✓ 1086ms | http |
| 167.103.144.127:8800 | ✓ 1148ms | 否 | ✓ 1246ms | ✓ 1482ms | ✓ 1375ms | http |
| 5.104.87.17:8051 | ✓ 1246ms | 否 | 否 | ✓ 1445ms | ✓ 1425ms | http |
| 120.92.108.86:7890 | ✓ 1191ms | 否 | ✓ 1715ms | ✓ 1602ms | ✓ 1376ms | http |
| 167.103.31.122:8800 | ✓ 1691ms | 否 | ✓ 1752ms | ✓ 1626ms | ✓ 1573ms | http |
| 45.167.124.52:8080 | ✓ 777ms | ✓ 1683ms | ✓ 932ms | 否 | 否 | http |
| 45.167.125.21:999 | ✓ 1053ms | ✓ 1954ms | ✓ 1382ms | ✓ 1754ms | ✓ 1505ms | http |
| 46.30.46.133:3128 | ✓ 1425ms | ✓ 1749ms | ✓ 757ms | 否 | 否 | http |
| 222.228.171.92:8080 | ✓ 1803ms | 否 | ✓ 662ms | 否 | ✓ 1919ms | http |
| 47.238.220.4:8888 | 否 | ✓ 1568ms | 否 | ✓ 1878ms | ✓ 800ms | http |
| 212.58.132.5:8888 | ✓ 1317ms | 否 | ✓ 1462ms | ✓ 1508ms | ✓ 1201ms | http |
| 59.46.216.131:30001 | ✓ 932ms | ✓ 1881ms | ✓ 964ms | 否 | 否 | http |
| 103.229.127.31:7890 | ✓ 1077ms | ✓ 1951ms | ✓ 1028ms | 否 | ✓ 1048ms | http |
| 8.219.195.129:1080 | ✓ 1290ms | ✓ 1819ms | ✓ 730ms | ✓ 1000ms | ✓ 794ms | http |
| 147.161.239.240:8800 | ✓ 736ms | ✓ 1875ms | ✓ 1016ms | ✓ 1692ms | ✓ 1288ms | http |
| 137.59.47.73:3128 | ✓ 1894ms | 否 | ✓ 1093ms | ✓ 1070ms | ✓ 1305ms | http |
| 149.56.24.51:3128 | ✓ 665ms | 否 | 否 | ✓ 1840ms | ✓ 1384ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1584ms | ✓ 1891ms | ✓ 1061ms | 否 | http |
| 114.237.77.231:1080 | 否 | ✓ 1066ms | ✓ 1964ms | ✓ 1882ms | ✓ 1946ms | http |
| 45.140.147.82:1081 | ✓ 787ms | ✓ 1748ms | ✓ 1479ms | ✓ 1791ms | ✓ 1404ms | http |
| 195.26.224.49:3128 | ✓ 894ms | 否 | ✓ 1085ms | 否 | ✓ 1808ms | http |
| 118.113.247.70:1080 | ✓ 1119ms | ✓ 1581ms | ✓ 1604ms | ✓ 1547ms | ✓ 1152ms | http |
| 121.130.199.80:3128 | ✓ 1936ms | ✓ 1653ms | 否 | ✓ 1310ms | ✓ 1102ms | http |
| 95.214.9.93:3128 | ✓ 1690ms | 否 | ✓ 1583ms | ✓ 1962ms | 否 | http |
| 223.16.170.103:80 | ✓ 1056ms | 否 | ✓ 1136ms | ✓ 1226ms | ✓ 1041ms | http |
| 5.104.87.17:8050 | ✓ 1736ms | 否 | ✓ 1312ms | ✓ 957ms | ✓ 668ms | http |
| 163.61.38.128:3128 | ✓ 1317ms | 否 | ✓ 1952ms | ✓ 1926ms | 否 | http |
| 163.61.38.126:3128 | ✓ 1040ms | 否 | ✓ 1296ms | ✓ 1710ms | ✓ 1765ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1002ms | ✓ 901ms | 否 | ✓ 891ms | http |
| 24.144.86.173:1080 | ✓ 291ms | ✓ 904ms | ✓ 238ms | ✓ 687ms | ✓ 484ms | http |
| 34.50.41.219:3128 | ✓ 657ms | 否 | ✓ 1061ms | ✓ 1948ms | ✓ 1081ms | http |
| 51.145.178.158:3128 | ✓ 1001ms | 否 | ✓ 1913ms | 否 | ✓ 1514ms | http |
| 116.80.63.178:3172 | ✓ 1892ms | 否 | ✓ 1868ms | ✓ 1842ms | ✓ 1710ms | http |
| 62.113.119.14:8080 | ✓ 802ms | ✓ 1584ms | ✓ 1208ms | ✓ 1681ms | ✓ 1389ms | http |
| 110.42.37.202:20005 | ✓ 1605ms | ✓ 1380ms | ✓ 1083ms | ✓ 1613ms | ✓ 1301ms | http |
| 5.196.101.18:3128 | ✓ 1403ms | 否 | ✓ 1427ms | 否 | ✓ 1701ms | http |
| 45.140.147.155:1082 | ✓ 618ms | ✓ 1808ms | ✓ 1146ms | 否 | ✓ 1547ms | http |
| 218.60.0.214:80 | ✓ 1268ms | 否 | ✓ 1573ms | ✓ 1585ms | ✓ 1278ms | http |
| 171.227.167.109:1008 | ✓ 1684ms | 否 | ✓ 1405ms | ✓ 1390ms | ✓ 1025ms | http |
| 103.163.80.170:1080 | 否 | 否 | ✓ 1219ms | ✓ 1361ms | ✓ 1363ms | http |
| 61.52.131.172:8443 | ✓ 867ms | ✓ 1191ms | ✓ 1034ms | ✓ 1175ms | ✓ 937ms | http |
| 103.39.51.207:8080 | ✓ 1392ms | 否 | 否 | ✓ 1740ms | ✓ 1497ms | http |
| 107.172.102.234:40621 | ✓ 1158ms | ✓ 605ms | ✓ 627ms | ✓ 704ms | ✓ 710ms | http |

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
