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

最后更新：2026-02-26 15:46:14 UTC（2026-02-26 23:46:14 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 101.47.73.135:3128 | ✓ 1112ms | 否 | ✓ 1105ms | ✓ 1302ms | ✓ 1358ms | http |
| 35.225.22.61:80 | ✓ 529ms | 否 | ✓ 899ms | ✓ 1183ms | ✓ 841ms | http |
| 211.230.49.122:3128 | ✓ 1035ms | ✓ 1755ms | 否 | ✓ 1337ms | ✓ 1200ms | http |
| 185.246.90.163:10808 | ✓ 1638ms | ✓ 1656ms | ✓ 1702ms | 否 | ✓ 1614ms | http |
| 72.56.59.17:61931 | ✓ 1538ms | 否 | ✓ 1645ms | 否 | ✓ 1907ms | http |
| 72.56.50.17:59787 | ✓ 1575ms | 否 | ✓ 1867ms | 否 | ✓ 1909ms | http |
| 103.84.95.54:7890 | ✓ 789ms | 否 | 否 | ✓ 1025ms | ✓ 822ms | http |
| 165.232.76.189:3128 | ✓ 815ms | 否 | 否 | ✓ 1922ms | ✓ 1363ms | http |
| 202.152.44.19:8081 | ✓ 1925ms | 否 | ✓ 1119ms | ✓ 1872ms | 否 | http |
| 14.56.107.244:3128 | ✓ 1755ms | ✓ 1759ms | ✓ 1449ms | ✓ 1679ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1113ms | ✓ 1313ms | ✓ 1642ms | 否 | ✓ 971ms | http |
| 202.152.44.18:8081 | ✓ 1950ms | 否 | ✓ 1122ms | ✓ 1932ms | 否 | http |
| 103.46.8.88:3125 | ✓ 1943ms | 否 | ✓ 1574ms | ✓ 1679ms | 否 | http |
| 72.56.59.56:63127 | ✓ 1496ms | 否 | ✓ 1613ms | 否 | ✓ 1907ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1538ms | ✓ 1286ms | ✓ 1641ms | ✓ 1586ms | http |
| 186.148.180.46:999 | ✓ 703ms | 否 | ✓ 533ms | 否 | ✓ 1305ms | http |
| 37.27.100.112:443 | ✓ 875ms | ✓ 1998ms | ✓ 1620ms | 否 | ✓ 1351ms | http |
| 120.92.212.16:8890 | ✓ 1072ms | ✓ 1678ms | ✓ 1094ms | 否 | 否 | http |
| 72.56.59.62:63133 | ✓ 1535ms | 否 | ✓ 1611ms | 否 | ✓ 1872ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1407ms | ✓ 1357ms | 否 | ✓ 1058ms | http |
| 72.56.59.23:61937 | ✓ 1551ms | 否 | ✓ 1587ms | 否 | ✓ 1837ms | http |
| 61.72.110.94:3128 | 否 | 否 | ✓ 819ms | ✓ 1210ms | ✓ 1496ms | http |
| 144.31.69.170:1080 | ✓ 1374ms | 否 | ✓ 1346ms | 否 | ✓ 1772ms | http |
| 103.113.70.189:1081 | 否 | ✓ 936ms | 否 | ✓ 1092ms | ✓ 895ms | http |
| 61.72.110.24:3128 | 否 | ✓ 1917ms | 否 | ✓ 1217ms | ✓ 1972ms | http |
| 36.147.78.166:443 | 否 | ✓ 1805ms | ✓ 1837ms | 否 | ✓ 1732ms | http |
| 36.147.78.166:80 | ✓ 1926ms | 否 | ✓ 1876ms | ✓ 1872ms | ✓ 1785ms | http |
| 120.46.152.136:3128 | ✓ 1717ms | 否 | ✓ 1972ms | ✓ 1963ms | ✓ 1278ms | http |
| 190.9.109.202:999 | ✓ 661ms | 否 | ✓ 1269ms | ✓ 1478ms | ✓ 1169ms | http |
| 168.235.110.63:3128 | ✓ 216ms | 否 | ✓ 1824ms | ✓ 1024ms | ✓ 776ms | http |
| 52.188.28.218:3128 | ✓ 247ms | 否 | ✓ 919ms | ✓ 1928ms | ✓ 722ms | http |
| 185.191.236.162:3128 | ✓ 1076ms | 否 | ✓ 875ms | ✓ 1589ms | ✓ 1144ms | http |
| 217.217.249.160:80 | ✓ 1793ms | 否 | ✓ 1355ms | 否 | ✓ 1425ms | http |
| 61.72.110.54:3128 | ✓ 1530ms | ✓ 1986ms | ✓ 1954ms | 否 | 否 | http |
| 35.234.17.221:8080 | ✓ 951ms | 否 | ✓ 1630ms | 否 | ✓ 1009ms | http |
| 116.80.48.38:7777 | ✓ 1860ms | 否 | 否 | ✓ 1921ms | ✓ 1739ms | http |
| 101.43.255.96:80 | ✓ 1978ms | 否 | ✓ 1130ms | ✓ 1783ms | 否 | http |
| 81.70.169.194:80 | ✓ 1485ms | ✓ 1538ms | ✓ 1512ms | 否 | ✓ 1060ms | http |
| 42.115.72.172:2087 | ✓ 1610ms | 否 | ✓ 1678ms | ✓ 1816ms | ✓ 1656ms | http |
| 38.183.146.25:8090 | ✓ 1504ms | 否 | ✓ 1681ms | ✓ 1587ms | ✓ 1566ms | http |
| 14.56.118.184:3128 | ✓ 1675ms | 否 | 否 | ✓ 1243ms | ✓ 1453ms | http |
| 14.56.118.164:3128 | ✓ 1667ms | 否 | 否 | ✓ 1210ms | ✓ 1014ms | http |
| 121.128.121.144:3128 | ✓ 1670ms | 否 | 否 | ✓ 1707ms | ✓ 909ms | http |
| 43.161.214.161:1081 | ✓ 1399ms | 否 | ✓ 1688ms | 否 | ✓ 1831ms | http |
| 121.128.121.244:3128 | ✓ 1701ms | 否 | 否 | ✓ 1222ms | ✓ 1175ms | http |
| 104.238.30.91:63900 | ✓ 1745ms | 否 | ✓ 1910ms | 否 | ✓ 1992ms | http |
| 61.72.221.244:3128 | ✓ 1749ms | 否 | 否 | ✓ 1514ms | ✓ 1497ms | http |
| 113.45.250.180:443 | ✓ 1105ms | 否 | ✓ 1180ms | 否 | ✓ 1086ms | http |
| 34.101.184.164:3128 | ✓ 1063ms | 否 | ✓ 1132ms | ✓ 1590ms | ✓ 1136ms | http |
| 89.251.9.11:3128 | ✓ 248ms | 否 | ✓ 101ms | ✓ 1535ms | ✓ 987ms | http |
| 8.219.97.248:80 | ✓ 1551ms | 否 | 否 | ✓ 1752ms | ✓ 1436ms | http |
| 94.177.131.12:3128 | ✓ 1032ms | 否 | ✓ 1221ms | ✓ 1002ms | ✓ 1075ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1025ms | 否 | ✓ 1150ms | ✓ 989ms | http |
| 81.177.48.54:2080 | ✓ 1319ms | 否 | ✓ 1842ms | ✓ 1906ms | ✓ 1752ms | http |
| 103.139.138.194:3128 | ✓ 1905ms | 否 | ✓ 1467ms | ✓ 1423ms | ✓ 1221ms | http |
| 162.240.154.26:3128 | ✓ 730ms | ✓ 1545ms | ✓ 1455ms | 否 | 否 | http |
| 152.32.255.24:27197 | ✓ 1950ms | 否 | ✓ 1870ms | ✓ 1494ms | ✓ 1311ms | http |
| 124.16.93.70:7890 | 否 | ✓ 1545ms | ✓ 1071ms | ✓ 1365ms | 否 | http |
| 223.113.134.102:22222 | ✓ 800ms | ✓ 1220ms | ✓ 802ms | 否 | 否 | http |
| 120.240.35.173:22222 | 否 | ✓ 1298ms | 否 | ✓ 1314ms | ✓ 985ms | http |
| 211.171.114.154:3128 | ✓ 1936ms | ✓ 1680ms | ✓ 1681ms | ✓ 1470ms | ✓ 1396ms | http |
| 223.113.134.104:22222 | ✓ 794ms | ✓ 1490ms | ✓ 1055ms | 否 | 否 | http |
| 120.232.242.119:22222 | ✓ 1233ms | ✓ 1316ms | ✓ 1075ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1624ms | ✓ 957ms | ✓ 680ms | ✓ 1062ms | ✓ 792ms | http |
| 113.59.32.162:22222 | ✓ 1327ms | ✓ 1960ms | 否 | ✓ 1410ms | 否 | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1628ms | ✓ 1977ms | ✓ 1614ms | http |
| 121.128.121.34:3128 | ✓ 1556ms | 否 | ✓ 1595ms | 否 | ✓ 1392ms | http |
| 175.215.54.252:3040 | ✓ 804ms | ✓ 1356ms | ✓ 1189ms | ✓ 1130ms | ✓ 929ms | http |
| 46.246.1.106:3128 | 否 | 否 | ✓ 1603ms | ✓ 1920ms | ✓ 1465ms | http |
| 217.217.254.94:8080 | 否 | 否 | ✓ 1207ms | ✓ 1506ms | ✓ 1462ms | http |
| 165.227.5.10:8888 | ✓ 736ms | 否 | 否 | ✓ 1887ms | ✓ 689ms | http |
| 192.71.213.85:9091 | ✓ 921ms | 否 | ✓ 1553ms | ✓ 1988ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1476ms | ✓ 1927ms | ✓ 1885ms | 否 | 否 | http |
| 71.204.156.52:443 | ✓ 985ms | 否 | ✓ 1046ms | ✓ 1378ms | 否 | http |
| 128.199.121.61:9090 | 否 | 否 | ✓ 1552ms | ✓ 1730ms | ✓ 1093ms | http |
| 170.246.183.250:3128 | ✓ 894ms | 否 | ✓ 776ms | 否 | ✓ 1808ms | http |
| 34.96.238.40:8080 | ✓ 1762ms | 否 | 否 | ✓ 1128ms | ✓ 993ms | http |

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
