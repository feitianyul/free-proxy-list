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

最后更新：2026-02-27 12:34:15 UTC（2026-02-27 20:34:15 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1271ms | ✓ 1947ms | ✓ 645ms | ✓ 1067ms | ✓ 808ms | http |
| 72.56.59.56:63127 | ✓ 1891ms | 否 | ✓ 1583ms | 否 | ✓ 1743ms | http |
| 81.177.48.54:2080 | ✓ 1563ms | 否 | ✓ 1730ms | 否 | ✓ 1699ms | http |
| 104.238.30.38:59741 | ✓ 1583ms | 否 | ✓ 1931ms | 否 | ✓ 1907ms | http |
| 72.56.59.62:63133 | ✓ 1890ms | 否 | ✓ 1584ms | 否 | ✓ 1742ms | http |
| 104.238.30.50:59741 | ✓ 1596ms | 否 | ✓ 1902ms | 否 | ✓ 1935ms | http |
| 104.238.30.39:59741 | ✓ 1579ms | 否 | ✓ 1934ms | 否 | ✓ 1907ms | http |
| 168.235.110.63:3128 | ✓ 519ms | ✓ 1324ms | ✓ 1053ms | ✓ 1158ms | ✓ 928ms | http |
| 34.142.0.1:10808 | ✓ 391ms | 否 | ✓ 1016ms | ✓ 1600ms | ✓ 1441ms | http |
| 147.45.216.148:1080 | ✓ 674ms | 否 | ✓ 1139ms | ✓ 1612ms | ✓ 1156ms | http |
| 120.92.212.16:8890 | ✓ 1154ms | ✓ 1455ms | 否 | ✓ 1513ms | 否 | http |
| 185.246.90.163:10808 | ✓ 609ms | 否 | 否 | ✓ 1975ms | ✓ 1312ms | http |
| 72.56.50.17:59787 | ✓ 1445ms | 否 | ✓ 1706ms | 否 | ✓ 1744ms | http |
| 104.238.30.91:63900 | ✓ 1562ms | 否 | ✓ 1647ms | 否 | ✓ 1903ms | http |
| 104.238.30.45:59741 | ✓ 1606ms | 否 | ✓ 1611ms | 否 | ✓ 1904ms | http |
| 104.238.30.86:63900 | ✓ 1598ms | 否 | ✓ 1647ms | 否 | ✓ 1902ms | http |
| 104.238.30.40:59741 | ✓ 1571ms | 否 | ✓ 1647ms | 否 | ✓ 1904ms | http |
| 35.234.17.221:8080 | 否 | 否 | ✓ 1251ms | ✓ 1224ms | ✓ 1030ms | http |
| 181.78.194.249:999 | 否 | 否 | ✓ 1614ms | ✓ 1768ms | ✓ 1497ms | http |
| 81.70.169.194:80 | ✓ 1144ms | 否 | ✓ 1241ms | 否 | ✓ 1130ms | http |
| 52.188.28.218:3128 | ✓ 79ms | ✓ 1253ms | ✓ 884ms | ✓ 1563ms | ✓ 1255ms | http |
| 35.225.22.61:80 | ✓ 854ms | 否 | ✓ 1091ms | ✓ 1156ms | ✓ 842ms | http |
| 94.177.131.12:3128 | ✓ 1495ms | 否 | ✓ 706ms | ✓ 1026ms | ✓ 864ms | http |
| 120.92.212.16:7890 | ✓ 1155ms | 否 | ✓ 1193ms | ✓ 1518ms | 否 | http |
| 14.56.107.244:3128 | 否 | ✓ 1986ms | ✓ 843ms | 否 | ✓ 974ms | http |
| 45.88.0.115:3128 | ✓ 1464ms | ✓ 1629ms | 否 | 否 | ✓ 1009ms | http |
| 45.88.0.114:3128 | ✓ 419ms | 否 | 否 | ✓ 1329ms | ✓ 1124ms | http |
| 45.88.0.113:3128 | ✓ 1726ms | 否 | ✓ 834ms | ✓ 1337ms | 否 | http |
| 45.88.0.117:3128 | ✓ 1716ms | 否 | ✓ 846ms | 否 | ✓ 1570ms | http |
| 107.155.65.87:13428 | ✓ 1167ms | 否 | ✓ 1797ms | ✓ 1385ms | 否 | http |
| 164.90.151.28:3128 | ✓ 1546ms | ✓ 1182ms | ✓ 478ms | ✓ 982ms | ✓ 1110ms | http |
| 104.238.30.63:63744 | ✓ 1608ms | 否 | ✓ 1683ms | 否 | ✓ 1903ms | http |
| 104.238.30.58:63744 | ✓ 1609ms | 否 | ✓ 1684ms | 否 | ✓ 1931ms | http |
| 45.88.0.111:3128 | ✓ 738ms | 否 | ✓ 736ms | 否 | ✓ 1582ms | http |
| 104.238.30.37:59741 | ✓ 1606ms | 否 | ✓ 1683ms | 否 | ✓ 1902ms | http |
| 37.187.109.70:10111 | ✓ 970ms | ✓ 1683ms | ✓ 757ms | ✓ 1832ms | 否 | http |
| 103.84.95.54:7890 | ✓ 1717ms | 否 | ✓ 862ms | ✓ 1081ms | ✓ 1094ms | http |
| 45.140.147.155:1081 | ✓ 467ms | ✓ 1773ms | ✓ 384ms | ✓ 1273ms | ✓ 915ms | http |
| 45.140.147.155:1082 | ✓ 470ms | ✓ 1688ms | ✓ 412ms | ✓ 1331ms | ✓ 923ms | http |
| 20.2.83.243:3128 | ✓ 886ms | 否 | ✓ 955ms | ✓ 1216ms | ✓ 903ms | http |
| 222.131.220.234:10808 | ✓ 1119ms | ✓ 1361ms | ✓ 1228ms | ✓ 1418ms | ✓ 1240ms | http |
| 101.43.255.96:80 | ✓ 1160ms | ✓ 1569ms | ✓ 1185ms | 否 | 否 | http |
| 185.243.218.43:49153 | ✓ 823ms | 否 | ✓ 1096ms | ✓ 1909ms | ✓ 1560ms | http |
| 195.123.209.48:3128 | ✓ 1284ms | 否 | ✓ 1135ms | ✓ 1799ms | ✓ 1804ms | http |
| 152.32.255.24:27197 | ✓ 1724ms | 否 | 否 | ✓ 1604ms | ✓ 1258ms | http |
| 59.46.216.131:30001 | ✓ 1153ms | ✓ 1698ms | ✓ 1986ms | 否 | ✓ 1289ms | http |
| 36.147.78.166:80 | 否 | ✓ 1902ms | ✓ 1982ms | ✓ 1968ms | ✓ 1882ms | http |
| 104.238.30.68:63744 | ✓ 1610ms | 否 | ✓ 1683ms | 否 | ✓ 1899ms | http |
| 62.113.119.14:8080 | ✓ 1013ms | 否 | ✓ 1353ms | 否 | ✓ 1394ms | http |
| 45.136.198.40:3128 | ✓ 658ms | ✓ 1358ms | ✓ 942ms | ✓ 1393ms | ✓ 1073ms | http |
| 138.124.53.25:7443 | ✓ 809ms | 否 | ✓ 1337ms | 否 | ✓ 1474ms | http |
| 103.215.36.88:17977 | ✓ 1245ms | 否 | ✓ 1325ms | ✓ 1655ms | ✓ 1340ms | http |
| 103.82.23.118:5178 | ✓ 1920ms | 否 | ✓ 1428ms | 否 | ✓ 1495ms | http |
| 210.223.44.230:3128 | ✓ 1905ms | ✓ 1529ms | ✓ 1247ms | ✓ 1971ms | ✓ 1032ms | http |
| 121.237.181.137:8888 | ✓ 1039ms | 否 | ✓ 1053ms | ✓ 1470ms | 否 | http |
| 120.46.152.136:3128 | ✓ 1287ms | ✓ 1613ms | ✓ 1623ms | ✓ 1555ms | ✓ 1299ms | http |
| 103.117.100.127:13082 | ✓ 837ms | 否 | ✓ 1211ms | ✓ 1064ms | ✓ 844ms | http |
| 103.125.181.135:9999 | ✓ 1753ms | 否 | 否 | ✓ 1672ms | ✓ 1176ms | http |
| 61.52.131.172:8443 | ✓ 1071ms | 否 | ✓ 1112ms | ✓ 1382ms | ✓ 1091ms | http |
| 101.47.73.135:3128 | ✓ 1397ms | 否 | 否 | ✓ 1849ms | ✓ 1121ms | http |
| 61.72.110.94:3128 | ✓ 1595ms | 否 | ✓ 1555ms | ✓ 1971ms | 否 | http |
| 43.248.11.162:1080 | ✓ 1566ms | 否 | ✓ 1399ms | ✓ 1901ms | ✓ 1692ms | http |
| 172.212.68.37:3128 | ✓ 348ms | ✓ 1962ms | ✓ 970ms | ✓ 1174ms | ✓ 733ms | http |
| 34.101.184.164:3128 | ✓ 1208ms | 否 | ✓ 1515ms | ✓ 1613ms | ✓ 1226ms | http |
| 103.113.70.189:1081 | 否 | ✓ 942ms | 否 | ✓ 1051ms | ✓ 789ms | http |

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
