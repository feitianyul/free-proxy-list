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

最后更新：2026-03-11 00:24:31 UTC（2026-03-11 08:24:31 UTC+8）

**代理总数：57**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.188:8443 | ✓ 1585ms | ✓ 901ms | ✓ 536ms | ✓ 983ms | ✓ 755ms | http |
| 45.136.130.191:8443 | ✓ 1581ms | ✓ 925ms | ✓ 515ms | ✓ 978ms | ✓ 763ms | http |
| 45.136.130.175:8443 | ✓ 1573ms | ✓ 1297ms | ✓ 325ms | ✓ 979ms | ✓ 754ms | http |
| 45.136.131.63:8443 | ✓ 1572ms | ✓ 1146ms | ✓ 834ms | ✓ 984ms | ✓ 753ms | http |
| 45.136.131.47:8443 | ✓ 1572ms | ✓ 1319ms | ✓ 832ms | ✓ 991ms | ✓ 756ms | http |
| 158.69.185.37:3129 | ✓ 1847ms | 否 | ✓ 1464ms | ✓ 1737ms | ✓ 1485ms | http |
| 194.213.18.200:443 | 否 | ✓ 1917ms | ✓ 1645ms | ✓ 1707ms | ✓ 1886ms | http |
| 1.231.81.166:3128 | ✓ 1945ms | ✓ 1189ms | ✓ 1133ms | 否 | ✓ 908ms | http |
| 217.77.102.18:3128 | ✓ 1250ms | ✓ 1875ms | ✓ 1489ms | ✓ 1766ms | ✓ 1380ms | http |
| 35.225.22.61:80 | ✓ 474ms | ✓ 1880ms | ✓ 958ms | ✓ 1278ms | ✓ 840ms | http |
| 178.236.245.17:3128 | ✓ 638ms | 否 | ✓ 1176ms | ✓ 1977ms | ✓ 1562ms | http |
| 178.236.245.59:3128 | ✓ 607ms | ✓ 1861ms | ✓ 1317ms | ✓ 1917ms | ✓ 1657ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1627ms | ✓ 1341ms | ✓ 1866ms | http |
| 178.156.224.42:3128 | ✓ 1278ms | 否 | ✓ 1552ms | ✓ 1873ms | 否 | http |
| 152.70.98.46:8888 | ✓ 1690ms | ✓ 1313ms | ✓ 1416ms | ✓ 1007ms | ✓ 822ms | http |
| 162.240.154.26:3128 | ✓ 810ms | ✓ 1066ms | ✓ 667ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1115ms | ✓ 1304ms | 否 | ✓ 1439ms | ✓ 1112ms | http |
| 202.155.12.161:443 | ✓ 1936ms | 否 | 否 | ✓ 1967ms | ✓ 1550ms | http |
| 190.9.109.198:999 | ✓ 771ms | ✓ 1168ms | ✓ 1121ms | ✓ 1214ms | ✓ 1007ms | http |
| 138.124.53.25:7443 | ✓ 608ms | 否 | 否 | ✓ 1624ms | ✓ 1138ms | http |
| 120.92.212.16:7890 | ✓ 1142ms | 否 | ✓ 1962ms | 否 | ✓ 1169ms | http |
| 81.70.169.194:80 | ✓ 1188ms | ✓ 1491ms | ✓ 1263ms | 否 | 否 | http |
| 101.43.255.96:80 | ✓ 1693ms | 否 | ✓ 1880ms | ✓ 1624ms | ✓ 1796ms | http |
| 185.115.74.185:8080 | ✓ 1130ms | ✓ 1532ms | ✓ 1419ms | 否 | 否 | http |
| 45.186.6.104:3128 | ✓ 1037ms | ✓ 1570ms | ✓ 1556ms | 否 | 否 | http |
| 46.183.25.8:443 | 否 | 否 | ✓ 1066ms | ✓ 1771ms | ✓ 1323ms | http |
| 120.92.212.16:8890 | ✓ 1150ms | 否 | ✓ 1223ms | ✓ 1410ms | ✓ 1739ms | http |
| 86.109.3.24:10012 | ✓ 310ms | ✓ 939ms | ✓ 517ms | ✓ 1070ms | ✓ 858ms | http |
| 86.109.3.24:9400 | ✓ 354ms | ✓ 954ms | ✓ 739ms | ✓ 1071ms | ✓ 853ms | http |
| 172.212.68.37:3128 | ✓ 363ms | ✓ 1662ms | ✓ 978ms | ✓ 1348ms | ✓ 1013ms | http |
| 138.124.90.140:1080 | ✓ 1383ms | ✓ 1846ms | ✓ 1257ms | ✓ 1756ms | ✓ 1233ms | http |
| 94.176.3.43:7443 | ✓ 1773ms | 否 | ✓ 831ms | ✓ 1788ms | ✓ 1738ms | http |
| 121.230.9.241:1080 | ✓ 1264ms | ✓ 1509ms | ✓ 1583ms | 否 | ✓ 1396ms | http |
| 121.230.9.33:1080 | ✓ 1529ms | 否 | ✓ 1312ms | ✓ 1707ms | 否 | http |
| 112.78.187.186:8080 | ✓ 1884ms | 否 | ✓ 1455ms | ✓ 1640ms | ✓ 1614ms | http |
| 114.4.251.26:8080 | ✓ 1883ms | 否 | ✓ 1498ms | ✓ 1699ms | ✓ 1658ms | http |
| 133.26.134.100:3128 | ✓ 1284ms | 否 | ✓ 1766ms | ✓ 1491ms | ✓ 982ms | http |
| 86.109.3.24:10007 | ✓ 1021ms | ✓ 1747ms | ✓ 523ms | ✓ 1093ms | ✓ 831ms | http |
| 86.109.3.24:10010 | ✓ 339ms | ✓ 945ms | ✓ 407ms | ✓ 1059ms | ✓ 856ms | http |
| 88.80.150.82:8080 | ✓ 1042ms | ✓ 1866ms | 否 | 否 | ✓ 1698ms | https |
| 91.107.141.42:8081 | ✓ 1203ms | 否 | ✓ 1813ms | ✓ 1715ms | ✓ 1637ms | http |
| 168.235.110.63:3128 | ✓ 805ms | ✓ 897ms | ✓ 892ms | ✓ 945ms | ✓ 735ms | http |
| 62.113.119.14:8080 | ✓ 546ms | ✓ 1464ms | ✓ 882ms | ✓ 1526ms | ✓ 1197ms | http |
| 37.139.33.145:1080 | ✓ 761ms | ✓ 1677ms | ✓ 1167ms | ✓ 1642ms | ✓ 1228ms | http |
| 34.101.184.164:3128 | ✓ 1903ms | 否 | ✓ 1840ms | ✓ 1593ms | ✓ 1605ms | http |
| 103.39.51.190:8080 | ✓ 1993ms | 否 | ✓ 1931ms | ✓ 1794ms | 否 | http |
| 45.136.198.40:3128 | ✓ 901ms | ✓ 1631ms | ✓ 685ms | 否 | ✓ 1621ms | http |
| 200.174.198.32:8888 | ✓ 1043ms | 否 | ✓ 1821ms | 否 | ✓ 1836ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1498ms | 否 | ✓ 1517ms | ✓ 1313ms | http |
| 152.42.213.210:8080 | ✓ 1839ms | 否 | 否 | ✓ 1949ms | ✓ 1042ms | http |
| 45.140.147.82:1081 | ✓ 989ms | ✓ 1988ms | ✓ 791ms | ✓ 1889ms | ✓ 1286ms | http |
| 101.47.73.135:3128 | ✓ 1105ms | 否 | 否 | ✓ 1595ms | ✓ 1537ms | http |
| 162.248.165.72:1080 | ✓ 1641ms | 否 | ✓ 910ms | 否 | ✓ 1998ms | http |
| 95.3.9.78:8080 | ✓ 760ms | 否 | ✓ 1470ms | ✓ 1625ms | ✓ 1240ms | http |
| 95.3.9.78:3128 | ✓ 1641ms | ✓ 1850ms | ✓ 1823ms | ✓ 1621ms | ✓ 1306ms | http |
| 61.52.131.172:8443 | ✓ 1094ms | ✓ 1364ms | ✓ 1069ms | ✓ 1351ms | ✓ 1069ms | http |
| 103.113.70.189:1081 | ✓ 209ms | ✓ 932ms | 否 | ✓ 1099ms | ✓ 1269ms | http |

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
