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

最后更新：2026-03-09 10:52:21 UTC（2026-03-09 18:52:21 UTC+8）

**代理总数：53**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 64.186.232.4:10808 | ✓ 553ms | ✓ 1111ms | ✓ 1669ms | ✓ 945ms | ✓ 1543ms | http |
| 101.47.73.135:3128 | ✓ 806ms | 否 | ✓ 823ms | ✓ 1073ms | ✓ 970ms | http |
| 202.155.12.161:443 | ✓ 1565ms | 否 | 否 | ✓ 1489ms | ✓ 1299ms | http |
| 47.77.193.180:1080 | ✓ 161ms | ✓ 811ms | ✓ 452ms | ✓ 751ms | ✓ 500ms | http |
| 1.231.81.166:3128 | ✓ 817ms | ✓ 1696ms | ✓ 1973ms | 否 | 否 | http |
| 116.80.49.172:3172 | ✓ 1512ms | 否 | 否 | ✓ 1756ms | ✓ 1647ms | http |
| 116.80.49.163:3172 | 否 | 否 | ✓ 1503ms | ✓ 1792ms | ✓ 1763ms | http |
| 119.18.145.241:20326 | 否 | 否 | ✓ 1633ms | ✓ 1976ms | ✓ 1717ms | http |
| 59.46.216.131:30001 | ✓ 839ms | ✓ 1122ms | ✓ 874ms | ✓ 1264ms | 否 | http |
| 116.80.49.169:3172 | ✓ 1499ms | 否 | 否 | ✓ 1866ms | ✓ 1879ms | http |
| 116.80.49.159:3172 | ✓ 1496ms | 否 | 否 | ✓ 1947ms | ✓ 1802ms | http |
| 116.80.49.167:3172 | ✓ 1575ms | 否 | ✓ 1909ms | ✓ 1981ms | ✓ 1809ms | http |
| 116.80.49.166:3172 | 否 | 否 | ✓ 1582ms | ✓ 1823ms | ✓ 1811ms | http |
| 116.80.49.165:3172 | ✓ 1498ms | 否 | ✓ 1986ms | ✓ 1971ms | ✓ 1952ms | http |
| 116.80.49.161:3172 | ✓ 1531ms | 否 | ✓ 1953ms | ✓ 1995ms | 否 | http |
| 121.237.181.137:8888 | ✓ 755ms | ✓ 1369ms | ✓ 1400ms | ✓ 1081ms | ✓ 1483ms | http |
| 115.231.181.40:8128 | ✓ 750ms | 否 | 否 | ✓ 1598ms | ✓ 665ms | http |
| 101.43.255.96:80 | 否 | ✓ 1012ms | ✓ 1778ms | ✓ 958ms | ✓ 821ms | http |
| 81.70.169.194:80 | ✓ 745ms | ✓ 1627ms | ✓ 1626ms | ✓ 928ms | ✓ 1549ms | http |
| 39.104.201.40:7890 | ✓ 1453ms | 否 | ✓ 1262ms | 否 | ✓ 965ms | http |
| 116.80.49.170:3172 | 否 | 否 | ✓ 1477ms | ✓ 1885ms | ✓ 1650ms | http |
| 165.227.5.10:8888 | ✓ 1238ms | ✓ 1522ms | 否 | 否 | ✓ 622ms | http |
| 120.92.212.16:8890 | ✓ 729ms | 否 | ✓ 1454ms | 否 | ✓ 910ms | http |
| 178.236.245.17:3128 | ✓ 1502ms | 否 | ✓ 1543ms | 否 | ✓ 1480ms | http |
| 120.92.212.16:7890 | ✓ 901ms | ✓ 1117ms | ✓ 912ms | ✓ 1157ms | ✓ 1764ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1237ms | ✓ 1734ms | ✓ 1875ms | ✓ 1698ms | http |
| 45.140.147.155:1081 | ✓ 624ms | 否 | ✓ 1218ms | 否 | ✓ 1257ms | http |
| 45.136.198.40:3128 | ✓ 913ms | ✓ 1976ms | ✓ 1923ms | 否 | 否 | http |
| 205.209.118.30:3138 | ✓ 314ms | 否 | ✓ 1595ms | 否 | ✓ 1016ms | http |
| 47.101.149.27:9010 | ✓ 1404ms | 否 | ✓ 1873ms | ✓ 1007ms | 否 | http |
| 67.169.98.211:443 | ✓ 1783ms | ✓ 1775ms | 否 | ✓ 1882ms | 否 | http |
| 152.42.213.210:8080 | ✓ 1573ms | 否 | ✓ 1267ms | ✓ 1592ms | ✓ 939ms | http |
| 194.213.18.200:443 | ✓ 1187ms | 否 | ✓ 742ms | 否 | ✓ 946ms | http |
| 168.235.110.63:3128 | ✓ 602ms | 否 | ✓ 1263ms | 否 | ✓ 963ms | http |
| 180.125.216.109:8118 | 否 | ✓ 950ms | ✓ 686ms | 否 | ✓ 771ms | http |
| 88.80.150.82:8080 | ✓ 1500ms | 否 | 否 | ✓ 1969ms | ✓ 1483ms | https |
| 116.80.49.168:3172 | ✓ 1471ms | 否 | ✓ 1621ms | 否 | ✓ 1650ms | http |
| 190.9.109.198:999 | ✓ 1092ms | ✓ 1582ms | ✓ 1395ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1416ms | 否 | 否 | ✓ 1827ms | ✓ 1569ms | http |
| 136.49.34.18:8888 | ✓ 1801ms | 否 | ✓ 1697ms | 否 | ✓ 1791ms | http |
| 178.236.245.59:3128 | ✓ 1321ms | ✓ 1953ms | ✓ 1791ms | 否 | 否 | http |
| 152.42.213.210:80 | 否 | 否 | ✓ 1426ms | ✓ 1627ms | ✓ 1492ms | http |
| 14.225.222.164:7890 | 否 | 否 | ✓ 869ms | ✓ 1905ms | ✓ 1736ms | http |
| 112.209.35.51:8082 | ✓ 1498ms | 否 | 否 | ✓ 1515ms | ✓ 1562ms | http |
| 35.225.22.61:80 | ✓ 419ms | ✓ 1354ms | 否 | ✓ 1143ms | 否 | http |
| 103.67.46.225:3125 | ✓ 1627ms | 否 | 否 | ✓ 1515ms | ✓ 1386ms | http |
| 113.132.112.110:9000 | ✓ 1833ms | 否 | ✓ 1343ms | ✓ 1445ms | 否 | http |
| 121.230.9.205:1080 | ✓ 1064ms | ✓ 1297ms | ✓ 1029ms | ✓ 1976ms | ✓ 1355ms | http |
| 82.146.45.118:3128 | ✓ 915ms | ✓ 1780ms | ✓ 1463ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1464ms | 否 | ✓ 1078ms | ✓ 916ms | ✓ 672ms | http |
| 61.72.110.54:3128 | ✓ 1482ms | 否 | ✓ 1447ms | 否 | ✓ 829ms | http |
| 144.208.127.181:3128 | ✓ 1519ms | 否 | ✓ 1347ms | 否 | ✓ 1504ms | http |
| 45.140.147.82:1081 | ✓ 1528ms | 否 | ✓ 1450ms | ✓ 1597ms | ✓ 1310ms | http |

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
